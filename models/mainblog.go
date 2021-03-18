// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Mainblog is an object representing the database table.
type Mainblog struct {
	ID         int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title      null.String `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	CreateTime null.Time   `boil:"createTime" json:"createTime,omitempty" toml:"createTime" yaml:"createTime,omitempty"`
	CreateUser null.String `boil:"createUser" json:"createUser,omitempty" toml:"createUser" yaml:"createUser,omitempty"`
	UpdateTime null.Time   `boil:"updateTime" json:"updateTime,omitempty" toml:"updateTime" yaml:"updateTime,omitempty"`
	UpdateUser null.String `boil:"updateUser" json:"updateUser,omitempty" toml:"updateUser" yaml:"updateUser,omitempty"`

	R *mainblogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mainblogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MainblogColumns = struct {
	ID         string
	Title      string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
}{
	ID:         "id",
	Title:      "title",
	CreateTime: "createTime",
	CreateUser: "createUser",
	UpdateTime: "updateTime",
	UpdateUser: "updateUser",
}

// Generated where

var MainblogWhere = struct {
	ID         whereHelperint
	Title      whereHelpernull_String
	CreateTime whereHelpernull_Time
	CreateUser whereHelpernull_String
	UpdateTime whereHelpernull_Time
	UpdateUser whereHelpernull_String
}{
	ID:         whereHelperint{field: "`mainblog`.`id`"},
	Title:      whereHelpernull_String{field: "`mainblog`.`title`"},
	CreateTime: whereHelpernull_Time{field: "`mainblog`.`createTime`"},
	CreateUser: whereHelpernull_String{field: "`mainblog`.`createUser`"},
	UpdateTime: whereHelpernull_Time{field: "`mainblog`.`updateTime`"},
	UpdateUser: whereHelpernull_String{field: "`mainblog`.`updateUser`"},
}

// MainblogRels is where relationship names are stored.
var MainblogRels = struct {
}{}

// mainblogR is where relationships are stored.
type mainblogR struct {
}

// NewStruct creates a new relationship struct
func (*mainblogR) NewStruct() *mainblogR {
	return &mainblogR{}
}

// mainblogL is where Load methods for each relationship are stored.
type mainblogL struct{}

var (
	mainblogAllColumns            = []string{"id", "title", "createTime", "createUser", "updateTime", "updateUser"}
	mainblogColumnsWithoutDefault = []string{"id", "title", "createTime", "createUser", "updateTime", "updateUser"}
	mainblogColumnsWithDefault    = []string{}
	mainblogPrimaryKeyColumns     = []string{"id"}
)

type (
	// MainblogSlice is an alias for a slice of pointers to Mainblog.
	// This should generally be used opposed to []Mainblog.
	MainblogSlice []*Mainblog
	// MainblogHook is the signature for custom Mainblog hook methods
	MainblogHook func(context.Context, boil.ContextExecutor, *Mainblog) error

	mainblogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mainblogType                 = reflect.TypeOf(&Mainblog{})
	mainblogMapping              = queries.MakeStructMapping(mainblogType)
	mainblogPrimaryKeyMapping, _ = queries.BindMapping(mainblogType, mainblogMapping, mainblogPrimaryKeyColumns)
	mainblogInsertCacheMut       sync.RWMutex
	mainblogInsertCache          = make(map[string]insertCache)
	mainblogUpdateCacheMut       sync.RWMutex
	mainblogUpdateCache          = make(map[string]updateCache)
	mainblogUpsertCacheMut       sync.RWMutex
	mainblogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var mainblogBeforeInsertHooks []MainblogHook
var mainblogBeforeUpdateHooks []MainblogHook
var mainblogBeforeDeleteHooks []MainblogHook
var mainblogBeforeUpsertHooks []MainblogHook

var mainblogAfterInsertHooks []MainblogHook
var mainblogAfterSelectHooks []MainblogHook
var mainblogAfterUpdateHooks []MainblogHook
var mainblogAfterDeleteHooks []MainblogHook
var mainblogAfterUpsertHooks []MainblogHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Mainblog) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mainblogBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Mainblog) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mainblogBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Mainblog) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mainblogBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Mainblog) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mainblogBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Mainblog) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mainblogAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Mainblog) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mainblogAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Mainblog) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mainblogAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Mainblog) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mainblogAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Mainblog) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mainblogAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMainblogHook registers your hook function for all future operations.
func AddMainblogHook(hookPoint boil.HookPoint, mainblogHook MainblogHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		mainblogBeforeInsertHooks = append(mainblogBeforeInsertHooks, mainblogHook)
	case boil.BeforeUpdateHook:
		mainblogBeforeUpdateHooks = append(mainblogBeforeUpdateHooks, mainblogHook)
	case boil.BeforeDeleteHook:
		mainblogBeforeDeleteHooks = append(mainblogBeforeDeleteHooks, mainblogHook)
	case boil.BeforeUpsertHook:
		mainblogBeforeUpsertHooks = append(mainblogBeforeUpsertHooks, mainblogHook)
	case boil.AfterInsertHook:
		mainblogAfterInsertHooks = append(mainblogAfterInsertHooks, mainblogHook)
	case boil.AfterSelectHook:
		mainblogAfterSelectHooks = append(mainblogAfterSelectHooks, mainblogHook)
	case boil.AfterUpdateHook:
		mainblogAfterUpdateHooks = append(mainblogAfterUpdateHooks, mainblogHook)
	case boil.AfterDeleteHook:
		mainblogAfterDeleteHooks = append(mainblogAfterDeleteHooks, mainblogHook)
	case boil.AfterUpsertHook:
		mainblogAfterUpsertHooks = append(mainblogAfterUpsertHooks, mainblogHook)
	}
}

// One returns a single mainblog record from the query.
func (q mainblogQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Mainblog, error) {
	o := &Mainblog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for mainblog")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Mainblog records from the query.
func (q mainblogQuery) All(ctx context.Context, exec boil.ContextExecutor) (MainblogSlice, error) {
	var o []*Mainblog

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Mainblog slice")
	}

	if len(mainblogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Mainblog records in the query.
func (q mainblogQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count mainblog rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mainblogQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if mainblog exists")
	}

	return count > 0, nil
}

// Mainblogs retrieves all the records using an executor.
func Mainblogs(mods ...qm.QueryMod) mainblogQuery {
	mods = append(mods, qm.From("`mainblog`"))
	return mainblogQuery{NewQuery(mods...)}
}

// FindMainblog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMainblog(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Mainblog, error) {
	mainblogObj := &Mainblog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `mainblog` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, mainblogObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from mainblog")
	}

	return mainblogObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Mainblog) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mainblog provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mainblogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mainblogInsertCacheMut.RLock()
	cache, cached := mainblogInsertCache[key]
	mainblogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mainblogAllColumns,
			mainblogColumnsWithDefault,
			mainblogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mainblogType, mainblogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mainblogType, mainblogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `mainblog` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `mainblog` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `mainblog` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, mainblogPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into mainblog")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for mainblog")
	}

CacheNoHooks:
	if !cached {
		mainblogInsertCacheMut.Lock()
		mainblogInsertCache[key] = cache
		mainblogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Mainblog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Mainblog) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	mainblogUpdateCacheMut.RLock()
	cache, cached := mainblogUpdateCache[key]
	mainblogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mainblogAllColumns,
			mainblogPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update mainblog, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `mainblog` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, mainblogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mainblogType, mainblogMapping, append(wl, mainblogPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update mainblog row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for mainblog")
	}

	if !cached {
		mainblogUpdateCacheMut.Lock()
		mainblogUpdateCache[key] = cache
		mainblogUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q mainblogQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for mainblog")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for mainblog")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MainblogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mainblogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `mainblog` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mainblogPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in mainblog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all mainblog")
	}
	return rowsAff, nil
}

var mySQLMainblogUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Mainblog) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mainblog provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mainblogColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMainblogUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	mainblogUpsertCacheMut.RLock()
	cache, cached := mainblogUpsertCache[key]
	mainblogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mainblogAllColumns,
			mainblogColumnsWithDefault,
			mainblogColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			mainblogAllColumns,
			mainblogPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert mainblog, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "mainblog", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `mainblog` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(mainblogType, mainblogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mainblogType, mainblogMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for mainblog")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(mainblogType, mainblogMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for mainblog")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for mainblog")
	}

CacheNoHooks:
	if !cached {
		mainblogUpsertCacheMut.Lock()
		mainblogUpsertCache[key] = cache
		mainblogUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Mainblog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Mainblog) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Mainblog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mainblogPrimaryKeyMapping)
	sql := "DELETE FROM `mainblog` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from mainblog")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for mainblog")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mainblogQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no mainblogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mainblog")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mainblog")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MainblogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(mainblogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mainblogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `mainblog` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mainblogPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mainblog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mainblog")
	}

	if len(mainblogAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Mainblog) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMainblog(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MainblogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MainblogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mainblogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `mainblog`.* FROM `mainblog` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mainblogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MainblogSlice")
	}

	*o = slice

	return nil
}

// MainblogExists checks if the Mainblog row exists.
func MainblogExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `mainblog` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if mainblog exists")
	}

	return exists, nil
}