package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bonavadeur/katyusha/pkg/bonalib"
)

var (
	mu sync.Mutex
)

func Get_Miporin_Matrix() ([][]int32, error) {
	bonalib.Log("[Get_Miporin_Matrix] Bắt đầu gọi tới MIPORIN_URL")
	resp, err := http.Get(MIPORIN_URL)
	if err != nil {
		bonalib.Log("[Get_Miporin_Matrix] Lỗi khi gọi GET:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bonalib.Log("[Get_Miporin_Matrix] Mã trạng thái HTTP không OK:", resp.StatusCode)
		return nil, fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		bonalib.Log("[Get_Miporin_Matrix] Lỗi khi đọc body:", err)
		return nil, err
	}

	var wait_response [][]int32
	err = json.Unmarshal(body, &wait_response)
	if err != nil {
		bonalib.Log("[Get_Miporin_Matrix] Lỗi khi parse JSON:", err)
		return nil, err
	}

	bonalib.Log("[Get_Miporin_Matrix] Đã lấy thành công ma trận")
	return wait_response, nil
}

func startPeriodicTask() {
	bonalib.Log("[startPeriodicTask] Bắt đầu task định kỳ mỗi 50s")
	ticker := time.NewTicker(50 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		bonalib.Log("[startPeriodicTask] Tick mới, gọi Get_Miporin_Matrix()")
		matrix, err := Get_Miporin_Matrix()
		if err != nil {
			bonalib.Log("[startPeriodicTask] Lỗi khi lấy ma trận:", err)
			continue
		}
		mu.Lock()
		MIPORIN_matrix = matrix
		mu.Unlock()
		bonalib.Log("[startPeriodicTask] Cập nhật MIPORIN_matrix thành công")
	}
}

func IP2Int32(ip string) string {
	var binaryIP string
	octets := strings.Split(ip, ".")

	if len(octets) != 4 {
		bonalib.Log("[IP2Int32] IP không hợp lệ (phải có 4 octet):", ip)
		return ""
	}

	for _, octet := range octets {
		if octet == "" {
			bonalib.Log("[IP2Int32] Octet rỗng trong IP:", ip)
			return ""
		}
		num, err := strconv.Atoi(octet)
		if err != nil {
			bonalib.Log("[IP2Int32] Lỗi khi chuyển đổi octet:", err)
			return ""
		}
		if num < 0 || num > 255 {
			bonalib.Log("[IP2Int32] Octet vượt giới hạn 0-255:", octet)
			return ""
		}
		binaryIP += fmt.Sprintf("%08b", num)
	}

	return binaryIP
}

func IsPodinPodcidr(ip string, cidr PodCIDR) bool {
	if ip == "" {
		bonalib.Log("[IsPodinPodcidr] IP đầu vào rỗng")
		return false
	}

	binaryIP := IP2Int32(ip)
	if binaryIP == "" {
		bonalib.Log("[IsPodinPodcidr] Không thể chuyển IP sang nhị phân:", ip)
		return false
	}

	binaryCIDR := IP2Int32(cidr.PodIPRange)
	if binaryCIDR == "" {
		bonalib.Log("[IsPodinPodcidr] Không thể chuyển PodIPRange sang nhị phân:", cidr.PodIPRange)
		return false
	}

	prefix := int(cidr.PodPrefix)
	if len(binaryIP) < prefix || len(binaryCIDR) < prefix {
		bonalib.Log("[IsPodinPodcidr] Độ dài không đủ:", "binaryIP:", len(binaryIP), "binaryCIDR:", len(binaryCIDR), "prefix:", prefix)
		return false
	}

	return binaryIP[:prefix] == binaryCIDR[:prefix]
}

func IPfromNode(ip string) string {
	if ip == "" {
		bonalib.Log("[IPfromNode] IP đầu vào rỗng")
		return "Request from Unknown Node"
	}

	for _, cidr := range PODCIDRS {
		if IsPodinPodcidr(ip, cidr) {
			bonalib.Log("[IPfromNode] IP", ip, "thuộc node", cidr.Nodename)
			return cidr.Nodename
		}
	}

	bonalib.Log("[IPfromNode] Không xác định được node cho IP:", ip)
	return "Request from Unknown Node"
}

func Choose(pool []int32) int {
	if len(pool) == 0 {
		bonalib.Log("[Choose] Cảnh báo: pool rỗng")
		return -1
	}

	ret := -1
	random := rand.Intn(100)
	for _, value := range pool {
		ret++
		if random < int(value) {
			bonalib.Log("[Choose] Chọn node index", ret, "với random =", random)
			return ret
		}
		random -= int(value)
	}

	bonalib.Log("[Choose] Không chọn được node nào (pool không đủ phân phối)")
	return -1
}
