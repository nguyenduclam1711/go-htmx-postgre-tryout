package repositories

import (
	"github.com/nguyenduclam1711/go-htmx-postgre-tryout/models"
)

type RepositoryFindConfig struct {
	Limit     int
	Skip      int
	Sort      [][2]string // something like: [][2]string{{"id", "asc"}}
	Condition []string    // simplified the conditions, I don't wanna create anoterh query builder: []string{"id = '1'", "country = 'Vietnam'"}
}

type CoreRepositoryer[M interface{}] interface {
	Create(m M) (M, error)
	// Find(c RepositoryFindConfig) ([]M, error)
	// FindOne(c RepositoryFindConfig) (M, error)
	// Update(c RepositoryFindConfig, m M) (numberOfUpdatedData int, err error)
	// Delete(c RepositoryFindConfig, m M) (numberOfDeletedData int, err error)
	// GetTableName() string
}

type CoreRepository[M interface{}] struct {
	ModelConfig models.ModelConfig
}

// TODO: Core repository create method
func (r *CoreRepository[M]) Create(v M) (M, error) {
	// currFields := []string{}
	// values := []string{}
	// insertedFields := []string{}
	// typeV := reflect.TypeOf(v)
	//
	// for i := 0; i < typeV.NumField(); i++ {
	// 	field := typeV.Field(i)
	// 	columnName := field.Tag.Get("db")
	// }
	// query := fmt.Sprintf(`
	// 	insert into %s (%s)
	// 	values (%s)
	// 	`, r.ModelConfig.TableName, "", "")
	return v, nil
}

func NewRepository[M interface{}](v M) CoreRepositoryer[M] {
	var res CoreRepositoryer[M] = &CoreRepository[M]{
		ModelConfig: models.MapModelConfigs[models.GetTableName(v)],
	}
	return res
}
