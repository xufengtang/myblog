// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMainblogs(t *testing.T) {
	t.Parallel()

	query := Mainblogs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMainblogsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Mainblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMainblogsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Mainblogs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Mainblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMainblogsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MainblogSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Mainblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMainblogsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MainblogExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Mainblog exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MainblogExists to return true, but got false.")
	}
}

func testMainblogsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mainblogFound, err := FindMainblog(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if mainblogFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMainblogsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Mainblogs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMainblogsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Mainblogs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMainblogsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mainblogOne := &Mainblog{}
	mainblogTwo := &Mainblog{}
	if err = randomize.Struct(seed, mainblogOne, mainblogDBTypes, false, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}
	if err = randomize.Struct(seed, mainblogTwo, mainblogDBTypes, false, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mainblogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mainblogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Mainblogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMainblogsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mainblogOne := &Mainblog{}
	mainblogTwo := &Mainblog{}
	if err = randomize.Struct(seed, mainblogOne, mainblogDBTypes, false, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}
	if err = randomize.Struct(seed, mainblogTwo, mainblogDBTypes, false, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mainblogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mainblogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Mainblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mainblogBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Mainblog) error {
	*o = Mainblog{}
	return nil
}

func mainblogAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Mainblog) error {
	*o = Mainblog{}
	return nil
}

func mainblogAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Mainblog) error {
	*o = Mainblog{}
	return nil
}

func mainblogBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Mainblog) error {
	*o = Mainblog{}
	return nil
}

func mainblogAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Mainblog) error {
	*o = Mainblog{}
	return nil
}

func mainblogBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Mainblog) error {
	*o = Mainblog{}
	return nil
}

func mainblogAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Mainblog) error {
	*o = Mainblog{}
	return nil
}

func mainblogBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Mainblog) error {
	*o = Mainblog{}
	return nil
}

func mainblogAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Mainblog) error {
	*o = Mainblog{}
	return nil
}

func testMainblogsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Mainblog{}
	o := &Mainblog{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mainblogDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Mainblog object: %s", err)
	}

	AddMainblogHook(boil.BeforeInsertHook, mainblogBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mainblogBeforeInsertHooks = []MainblogHook{}

	AddMainblogHook(boil.AfterInsertHook, mainblogAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mainblogAfterInsertHooks = []MainblogHook{}

	AddMainblogHook(boil.AfterSelectHook, mainblogAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mainblogAfterSelectHooks = []MainblogHook{}

	AddMainblogHook(boil.BeforeUpdateHook, mainblogBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mainblogBeforeUpdateHooks = []MainblogHook{}

	AddMainblogHook(boil.AfterUpdateHook, mainblogAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mainblogAfterUpdateHooks = []MainblogHook{}

	AddMainblogHook(boil.BeforeDeleteHook, mainblogBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mainblogBeforeDeleteHooks = []MainblogHook{}

	AddMainblogHook(boil.AfterDeleteHook, mainblogAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mainblogAfterDeleteHooks = []MainblogHook{}

	AddMainblogHook(boil.BeforeUpsertHook, mainblogBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mainblogBeforeUpsertHooks = []MainblogHook{}

	AddMainblogHook(boil.AfterUpsertHook, mainblogAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mainblogAfterUpsertHooks = []MainblogHook{}
}

func testMainblogsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Mainblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMainblogsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mainblogColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Mainblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMainblogsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMainblogsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MainblogSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMainblogsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Mainblogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mainblogDBTypes = map[string]string{`ID`: `int`, `Title`: `varchar`, `CreateTime`: `datetime`, `CreateUser`: `varchar`, `UpdateTime`: `datetime`, `UpdateUser`: `varchar`}
	_               = bytes.MinRead
)

func testMainblogsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mainblogPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mainblogAllColumns) == len(mainblogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Mainblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMainblogsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mainblogAllColumns) == len(mainblogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Mainblog{}
	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Mainblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mainblogDBTypes, true, mainblogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mainblogAllColumns, mainblogPrimaryKeyColumns) {
		fields = mainblogAllColumns
	} else {
		fields = strmangle.SetComplement(
			mainblogAllColumns,
			mainblogPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := MainblogSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMainblogsUpsert(t *testing.T) {
	t.Parallel()

	if len(mainblogAllColumns) == len(mainblogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLMainblogUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Mainblog{}
	if err = randomize.Struct(seed, &o, mainblogDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Mainblog: %s", err)
	}

	count, err := Mainblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mainblogDBTypes, false, mainblogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Mainblog struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Mainblog: %s", err)
	}

	count, err = Mainblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
