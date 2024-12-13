package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing
import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"fmt"

	"github.com/bonavadeur/katyusha/pkg/bonalib"
)

var (
	mu sync.Mutex
)

func Get_Miporin_Matrix() ([][]int32, error) {
	resp, err := http.Get(MIPORIN_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Kiểm tra mã trạng thái HTTP
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var wait_response [][]int32
	err = json.Unmarshal(body, &wait_response)
	if err != nil {
		return nil, err
	}

	return wait_response, nil
}
func startPeriodicTask() {
	// Tạo ticker mỗi 2 giây
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop() // Đảm bảo ticker được dừng khi kết thúc chương trình

	for range ticker.C {
		// Mỗi 2 giây, gọi hàm getMatrix()
		matrix, err := Get_Miporin_Matrix()
		if err != nil {
			bonalib.Log("Lỗi khi lấy ma trận:", err)
			continue
		}
		mu.Lock()
		MIPORIN_matrix = matrix
		mu.Unlock()

		// bonalib.Log("Ma trận Weight được cập nhật:", MIPORIN_matrix)
	}

}

// Hàm chuyển đổi IP từ dạng chuỗi 4octet thập phân sang dạng nhị phân 32 bit
func IP2Int32(ip string) string {
	var binaryIP string
	// Tách chuỗi IP thành các octet
	octets := strings.Split(ip, ".")

	for _, octet := range octets {
		num, err := strconv.Atoi(octet)
		if err != nil {
			bonalib.Log("Lỗi khi chuyển đổi:", err)
			return ""
		}
		// Chuyển số nguyên sang dạng nhị phân và đảm bảo có 8 bit
		binary := fmt.Sprintf("%08b", num)
		binaryIP += binary // Nối nhị phân vào chuỗi kết quả
	}

	return binaryIP
}

func IsPodinPodcidr(ip string, cidr PodCIDR) bool {
	// Chuyển đổi IP sang dạng nhị phân
	binaryIP := IP2Int32(ip)
	// Chuyển đổi CIDR sang dạng nhị phân
	binaryCIDR := IP2Int32(cidr.PodIPRange)
	return binaryIP[:cidr.PodPrefix] == binaryCIDR[:cidr.PodPrefix]
}

func IPfromNode(ip string) string {
	for _ , cidr := range PODCIDRS {
		if IsPodinPodcidr(ip, cidr) {
			// bonalib.Info("Request from node :", cidr.Nodename)
			return cidr.Nodename
		}
	}
	return "Request from Unknown Node"
}

func gachaNodeTarget(random int,Fight []int32) int {
	var ret int = -1;
	for _ , value := range Fight {
		ret++
		if random < int(value) {
			return ret // trả về node


		}
		random -= int(value)
	}
	return -1
}