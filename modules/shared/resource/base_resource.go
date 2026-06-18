package resource

import (
	"gin-money-manager-api/modules/shared/dto"
	"gin-money-manager-api/modules/shared/helper"
	"gin-money-manager-api/modules/shared/repository"
	"gin-money-manager-api/modules/shared/repository/options"
	"gin-money-manager-api/modules/shared/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type BaseResource[
	Entity any,
	CreateDto any,
	UpdateDto any,
] struct {
	Repository repository.BaseRepository[Entity]

	Relationships []string
	SearchFields  []string
	Filter        string

	Paginate bool
	Limit    int
	Page     int
}

func (r *BaseResource[E, C, U]) Index(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))
	if err != nil {
		limit = 50
	}

	var findAllOptions = &options.FindAllOptions{
		Relationships: r.Relationships,
		Search: &options.SearchOptions{
			Keyword: c.Query("search"),
			Fields:  r.SearchFields,
		},
		Paginate: r.Paginate,
		Limit:    limit,
		Page:     page,
	}

	if r.Filter != "" {
		findAllOptions.Filters = []options.Filter{
			{
				Field:    r.Filter,
				Operator: "=",
				Value:    c.Param(r.Filter),
			},
		}
	}

	entities, total, err := r.Repository.FindAll(findAllOptions)

	if r.Paginate {
		pagination := dto.Pagination{
			Data:  entities,
			Page:  page,
			Limit: limit,
			Total: total,
		}

		response.Success(c, pagination, err)
	} else {
		response.Success(c, entities, err)
	}

}

func (r *BaseResource[E, C, U]) Show(c *gin.Context) {
	id := c.Param("id")

	entity, err := r.Repository.Find(id, nil)

	response.Success(c, entity, err)
}

func (r *BaseResource[E, C, U]) Create(c *gin.Context) {
	var body C

	if !helper.BindAndValidate(c, &body) {
		return
	}

	var entity E

	copier.Copy(&entity, &body)

	result, err := r.Repository.Create(entity)

	response.Created(c, result, err)
}

func (r *BaseResource[E, C, U]) Update(c *gin.Context) {
	id := c.Param("id")

	var body U

	if !helper.BindAndValidate(c, &body) {
		return
	}

	entity, err := r.Repository.Find(id, nil)

	if err != nil {
		response.NotFound(c)
		return
	}

	copier.Copy(&entity, &body)

	entity, err = r.Repository.Update(entity)

	response.Success(c, entity, err)
}

func (r *BaseResource[E, C, U]) Delete(c *gin.Context) {
	id := c.Param("id")

	entity, err := r.Repository.Find(id, nil)

	if err != nil {
		response.NotFound(c)
		return
	}

	err = r.Repository.Delete(id)

	response.Success(c, entity, err)
}
