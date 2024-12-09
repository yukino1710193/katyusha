package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing
import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"

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
