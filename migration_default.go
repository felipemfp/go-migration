package migration

import "time"

type BaseMigration struct {
	id          time.Time
	description string
}

// GetID returns the ID of the migration.
func (m *BaseMigration) GetID() time.Time{
	return m.id
}

// GetDescription returns the ID of the migration.
func (m *BaseMigration) GetDescription() string {
	return m.description
}

// DefaultMigration is the default implementation of the migration.Migration.
//
// It is designed to provide a coded implementaiton of a migration. It receives
// an up and down anonymous methods to be ran while executing the migration.
//
// This implementation is used by the migration.CodeSource implemenation of the
// migration.Source.
type DefaultMigration struct {
	BaseMigration
	do      Handler
	undo    Handler
	manager Manager
}

// Handler is the signature of the up and down methods that a migration
// will receive.
type Handler func() error

// NewMigration returns a new instance of migration.Migration with all the
// required properties initialized.
//
// If a handler is provided it will assigned to the Up method. If a second is
// provided, it will be assigned to the Down method.
func NewMigration(id time.Time, description string, handlers ...Handler) *DefaultMigration {
	var do, undo Handler
	if len(handlers) > 0 {
		do = handlers[0]
	}
	if len(handlers) > 1 {
		undo = handlers[1]
	}
	return &DefaultMigration{
		BaseMigration: BaseMigration{
			id:          id,
			description: description,
		},
		do:   do,
		undo: undo,
	}
}

// Up calls the up action of the migration.
func (m *DefaultMigration) Do() error {
	return m.do()
}

// Down calls the down action of the migration.
func (m *DefaultMigration) Undo() error {
	return m.undo()
}

// GetManager returns the reference of the manager that is executing the
// migration.
func (m *DefaultMigration) GetManager() Manager {
	return m.manager
}

// SetManager set the reference of the manager that is executing the migration.
//
// It returns itself for sugar syntax.
func (m *DefaultMigration) SetManager(manager Manager) Migration {
	m.manager = manager
	return m
}