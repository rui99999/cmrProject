package cmrProj
import (
	"fmt"
	"github.com/go-sql-driver/mysql" 
)

func CmrFunc(string2 string)  {
	fmt.Printf(string2+":"+mysql.ErrPktSync.Error())
}