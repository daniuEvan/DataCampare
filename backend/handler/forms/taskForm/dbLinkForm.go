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
	ID         uint   `forms:"ID" json:"ID" binding:"-"`
	LinkName   string `forms:"LinkName" json:"LinkName" binding:"required"`
	DBType     string `forms:"DBType" json:"DBType" binding:"required,dbType"`
	DBHost     string `forms:"DBHost" json:"DBHost" binding:"required"`
	DBPort     uint   `forms:"DBPort" json:"DBPort" binding:"required,gt=0,lte=65535"`
	DBName     string `forms:"DBName" json:"DBName" binding:"required"`
	DBUsername string `forms:"DBUsername" json:"DBUsername" binding:"required"`
	DBPassword string `forms:"DBPassword" json:"DBPassword" binding:"required,gte=6,lt=18"`
}
