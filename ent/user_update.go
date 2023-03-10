// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"realworld/ent/article"
	"realworld/ent/predicate"
	"realworld/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetLevel sets the "level" field.
func (uu *UserUpdate) SetLevel(u uint32) *UserUpdate {
	uu.mutation.ResetLevel()
	uu.mutation.SetLevel(u)
	return uu
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLevel(u *uint32) *UserUpdate {
	if u != nil {
		uu.SetLevel(*u)
	}
	return uu
}

// AddLevel adds u to the "level" field.
func (uu *UserUpdate) AddLevel(u int32) *UserUpdate {
	uu.mutation.AddLevel(u)
	return uu
}

// SetBio sets the "bio" field.
func (uu *UserUpdate) SetBio(s string) *UserUpdate {
	uu.mutation.SetBio(s)
	return uu
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBio(s *string) *UserUpdate {
	if s != nil {
		uu.SetBio(*s)
	}
	return uu
}

// ClearBio clears the value of the "bio" field.
func (uu *UserUpdate) ClearBio() *UserUpdate {
	uu.mutation.ClearBio()
	return uu
}

// SetImage sets the "image" field.
func (uu *UserUpdate) SetImage(s string) *UserUpdate {
	uu.mutation.SetImage(s)
	return uu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (uu *UserUpdate) SetNillableImage(s *string) *UserUpdate {
	if s != nil {
		uu.SetImage(*s)
	}
	return uu
}

// ClearImage clears the value of the "image" field.
func (uu *UserUpdate) ClearImage() *UserUpdate {
	uu.mutation.ClearImage()
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetPoints sets the "points" field.
func (uu *UserUpdate) SetPoints(u uint32) *UserUpdate {
	uu.mutation.ResetPoints()
	uu.mutation.SetPoints(u)
	return uu
}

// SetNillablePoints sets the "points" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePoints(u *uint32) *UserUpdate {
	if u != nil {
		uu.SetPoints(*u)
	}
	return uu
}

// AddPoints adds u to the "points" field.
func (uu *UserUpdate) AddPoints(u int32) *UserUpdate {
	uu.mutation.AddPoints(u)
	return uu
}

// AddArticleIDs adds the "articles" edge to the Article entity by IDs.
func (uu *UserUpdate) AddArticleIDs(ids ...int) *UserUpdate {
	uu.mutation.AddArticleIDs(ids...)
	return uu
}

// AddArticles adds the "articles" edges to the Article entity.
func (uu *UserUpdate) AddArticles(a ...*Article) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddArticleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearArticles clears all "articles" edges to the Article entity.
func (uu *UserUpdate) ClearArticles() *UserUpdate {
	uu.mutation.ClearArticles()
	return uu
}

// RemoveArticleIDs removes the "articles" edge to Article entities by IDs.
func (uu *UserUpdate) RemoveArticleIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveArticleIDs(ids...)
	return uu
}

// RemoveArticles removes "articles" edges to Article entities.
func (uu *UserUpdate) RemoveArticles(a ...*Article) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveArticleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Level(); ok {
		_spec.SetField(user.FieldLevel, field.TypeUint32, value)
	}
	if value, ok := uu.mutation.AddedLevel(); ok {
		_spec.AddField(user.FieldLevel, field.TypeUint32, value)
	}
	if value, ok := uu.mutation.Bio(); ok {
		_spec.SetField(user.FieldBio, field.TypeString, value)
	}
	if uu.mutation.BioCleared() {
		_spec.ClearField(user.FieldBio, field.TypeString)
	}
	if value, ok := uu.mutation.Image(); ok {
		_spec.SetField(user.FieldImage, field.TypeString, value)
	}
	if uu.mutation.ImageCleared() {
		_spec.ClearField(user.FieldImage, field.TypeString)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Points(); ok {
		_spec.SetField(user.FieldPoints, field.TypeUint32, value)
	}
	if value, ok := uu.mutation.AddedPoints(); ok {
		_spec.AddField(user.FieldPoints, field.TypeUint32, value)
	}
	if uu.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ArticlesTable,
			Columns: []string{user.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedArticlesIDs(); len(nodes) > 0 && !uu.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ArticlesTable,
			Columns: []string{user.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ArticlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ArticlesTable,
			Columns: []string{user.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetLevel sets the "level" field.
func (uuo *UserUpdateOne) SetLevel(u uint32) *UserUpdateOne {
	uuo.mutation.ResetLevel()
	uuo.mutation.SetLevel(u)
	return uuo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLevel(u *uint32) *UserUpdateOne {
	if u != nil {
		uuo.SetLevel(*u)
	}
	return uuo
}

// AddLevel adds u to the "level" field.
func (uuo *UserUpdateOne) AddLevel(u int32) *UserUpdateOne {
	uuo.mutation.AddLevel(u)
	return uuo
}

// SetBio sets the "bio" field.
func (uuo *UserUpdateOne) SetBio(s string) *UserUpdateOne {
	uuo.mutation.SetBio(s)
	return uuo
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBio(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetBio(*s)
	}
	return uuo
}

// ClearBio clears the value of the "bio" field.
func (uuo *UserUpdateOne) ClearBio() *UserUpdateOne {
	uuo.mutation.ClearBio()
	return uuo
}

// SetImage sets the "image" field.
func (uuo *UserUpdateOne) SetImage(s string) *UserUpdateOne {
	uuo.mutation.SetImage(s)
	return uuo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableImage(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetImage(*s)
	}
	return uuo
}

// ClearImage clears the value of the "image" field.
func (uuo *UserUpdateOne) ClearImage() *UserUpdateOne {
	uuo.mutation.ClearImage()
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetPoints sets the "points" field.
func (uuo *UserUpdateOne) SetPoints(u uint32) *UserUpdateOne {
	uuo.mutation.ResetPoints()
	uuo.mutation.SetPoints(u)
	return uuo
}

// SetNillablePoints sets the "points" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePoints(u *uint32) *UserUpdateOne {
	if u != nil {
		uuo.SetPoints(*u)
	}
	return uuo
}

// AddPoints adds u to the "points" field.
func (uuo *UserUpdateOne) AddPoints(u int32) *UserUpdateOne {
	uuo.mutation.AddPoints(u)
	return uuo
}

// AddArticleIDs adds the "articles" edge to the Article entity by IDs.
func (uuo *UserUpdateOne) AddArticleIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddArticleIDs(ids...)
	return uuo
}

// AddArticles adds the "articles" edges to the Article entity.
func (uuo *UserUpdateOne) AddArticles(a ...*Article) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddArticleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearArticles clears all "articles" edges to the Article entity.
func (uuo *UserUpdateOne) ClearArticles() *UserUpdateOne {
	uuo.mutation.ClearArticles()
	return uuo
}

// RemoveArticleIDs removes the "articles" edge to Article entities by IDs.
func (uuo *UserUpdateOne) RemoveArticleIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveArticleIDs(ids...)
	return uuo
}

// RemoveArticles removes "articles" edges to Article entities.
func (uuo *UserUpdateOne) RemoveArticles(a ...*Article) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveArticleIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Level(); ok {
		_spec.SetField(user.FieldLevel, field.TypeUint32, value)
	}
	if value, ok := uuo.mutation.AddedLevel(); ok {
		_spec.AddField(user.FieldLevel, field.TypeUint32, value)
	}
	if value, ok := uuo.mutation.Bio(); ok {
		_spec.SetField(user.FieldBio, field.TypeString, value)
	}
	if uuo.mutation.BioCleared() {
		_spec.ClearField(user.FieldBio, field.TypeString)
	}
	if value, ok := uuo.mutation.Image(); ok {
		_spec.SetField(user.FieldImage, field.TypeString, value)
	}
	if uuo.mutation.ImageCleared() {
		_spec.ClearField(user.FieldImage, field.TypeString)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Points(); ok {
		_spec.SetField(user.FieldPoints, field.TypeUint32, value)
	}
	if value, ok := uuo.mutation.AddedPoints(); ok {
		_spec.AddField(user.FieldPoints, field.TypeUint32, value)
	}
	if uuo.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ArticlesTable,
			Columns: []string{user.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedArticlesIDs(); len(nodes) > 0 && !uuo.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ArticlesTable,
			Columns: []string{user.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ArticlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ArticlesTable,
			Columns: []string{user.ArticlesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
