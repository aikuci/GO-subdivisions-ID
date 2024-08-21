package mapper

type CruderMapper[TEntity any, TModel any] interface {
	ModelToResponse(entity *TEntity) *TModel
}
