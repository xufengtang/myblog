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

func testSmallblogs(t *testing.T) {
	t.Parallel()

	query := Smallblogs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSmallblogsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
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

	count, err := Smallblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSmallblogsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Smallblogs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Smallblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSmallblogsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SmallblogSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Smallblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSmallblogsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SmallblogExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Smallblog exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SmallblogExists to return true, but got false.")
	}
}

func testSmallblogsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	smallblogFound, err := FindSmallblog(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if smallblogFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSmallblogsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Smallblogs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSmallblogsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Smallblogs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSmallblogsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	smallblogOne := &Smallblog{}
	smallblogTwo := &Smallblog{}
	if err = randomize.Struct(seed, smallblogOne, smallblogDBTypes, false, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}
	if err = randomize.Struct(seed, smallblogTwo, smallblogDBTypes, false, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = smallblogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = smallblogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Smallblogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSmallblogsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	smallblogOne := &Smallblog{}
	smallblogTwo := &Smallblog{}
	if err = randomize.Struct(seed, smallblogOne, smallblogDBTypes, false, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}
	if err = randomize.Struct(seed, smallblogTwo, smallblogDBTypes, false, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = smallblogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = smallblogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Smallblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func smallblogBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Smallblog) error {
	*o = Smallblog{}
	return nil
}

func smallblogAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Smallblog) error {
	*o = Smallblog{}
	return nil
}

func smallblogAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Smallblog) error {
	*o = Smallblog{}
	return nil
}

func smallblogBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Smallblog) error {
	*o = Smallblog{}
	return nil
}

func smallblogAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Smallblog) error {
	*o = Smallblog{}
	return nil
}

func smallblogBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Smallblog) error {
	*o = Smallblog{}
	return nil
}

func smallblogAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Smallblog) error {
	*o = Smallblog{}
	return nil
}

func smallblogBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Smallblog) error {
	*o = Smallblog{}
	return nil
}

func smallblogAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Smallblog) error {
	*o = Smallblog{}
	return nil
}

func testSmallblogsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Smallblog{}
	o := &Smallblog{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, smallblogDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Smallblog object: %s", err)
	}

	AddSmallblogHook(boil.BeforeInsertHook, smallblogBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	smallblogBeforeInsertHooks = []SmallblogHook{}

	AddSmallblogHook(boil.AfterInsertHook, smallblogAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	smallblogAfterInsertHooks = []SmallblogHook{}

	AddSmallblogHook(boil.AfterSelectHook, smallblogAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	smallblogAfterSelectHooks = []SmallblogHook{}

	AddSmallblogHook(boil.BeforeUpdateHook, smallblogBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	smallblogBeforeUpdateHooks = []SmallblogHook{}

	AddSmallblogHook(boil.AfterUpdateHook, smallblogAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	smallblogAfterUpdateHooks = []SmallblogHook{}

	AddSmallblogHook(boil.BeforeDeleteHook, smallblogBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	smallblogBeforeDeleteHooks = []SmallblogHook{}

	AddSmallblogHook(boil.AfterDeleteHook, smallblogAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	smallblogAfterDeleteHooks = []SmallblogHook{}

	AddSmallblogHook(boil.BeforeUpsertHook, smallblogBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	smallblogBeforeUpsertHooks = []SmallblogHook{}

	AddSmallblogHook(boil.AfterUpsertHook, smallblogAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	smallblogAfterUpsertHooks = []SmallblogHook{}
}

func testSmallblogsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Smallblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSmallblogsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(smallblogColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Smallblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSmallblogsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
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

func testSmallblogsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SmallblogSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSmallblogsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Smallblogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	smallblogDBTypes = map[string]string{`ID`: `int`, `ParentId`: `int`, `Title`: `varchar`, `Content`: `varchar`, `CreateTime`: `datetime`, `CreateUser`: `varchar`, `UpdateTime`: `datetime`, `UpdateUser`: `varchar`}
	_                = bytes.MinRead
)

func testSmallblogsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(smallblogPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(smallblogAllColumns) == len(smallblogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Smallblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSmallblogsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(smallblogAllColumns) == len(smallblogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Smallblog{}
	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Smallblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, smallblogDBTypes, true, smallblogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(smallblogAllColumns, smallblogPrimaryKeyColumns) {
		fields = smallblogAllColumns
	} else {
		fields = strmangle.SetComplement(
			smallblogAllColumns,
			smallblogPrimaryKeyColumns,
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

	slice := SmallblogSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSmallblogsUpsert(t *testing.T) {
	t.Parallel()

	if len(smallblogAllColumns) == len(smallblogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLSmallblogUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Smallblog{}
	if err = randomize.Struct(seed, &o, smallblogDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Smallblog: %s", err)
	}

	count, err := Smallblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, smallblogDBTypes, false, smallblogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Smallblog struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Smallblog: %s", err)
	}

	count, err = Smallblogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
