/**
 * @date: 2022/3/3
 * @desc: ...
 */

package taskForm

//
// DBLinkForm
// @Description: 数据连接表单
//
type DBLinkForm struct {
	LinkName   string `forms:"LinkName" json:"LinkName" binding:""`
	DBType     string `forms:"DBType" json:"DBType" binding:""`
	DBHost     string `forms:"DBHost" json:"DBHost" binding:""`
	DBPort     uint   `forms:"DBPort" json:"DBPort" binding:"gt=0"`
	DBName     string `forms:"DBName" json:"DBName" binding:""`
	DBUsername string `forms:"DBUsername" json:"DBUsername" binding:""`
	DBPassword string `forms:"DBPassword" json:"DBPassword" binding:""`
}
