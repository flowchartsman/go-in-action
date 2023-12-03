package productdb

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/flowchartsman/go-in-action/catco/catsql"
)

// ErrItemNotFound is returned if the user attempts to retrieve or update a
// non-existing item in the product database
var ErrItemNotFound = errors.New("item not found")

// Client represents a connection to the product database.
type Client struct {
	dbhost string
	db     *catsql.DB
}

// NewClient creates a new client, ready to connect to dbhost.
// Note: okay, not really. The book examples actually use a temporary SQLite
// database, but many of the database libraries you will use will be connecting
// to some remote host, so I thought it prudent to simulate it here.
func NewClient(dbhost string) *Client {
	return &Client{
		dbhost: dbhost,
	}
}

// Connect connects to the backing database. In the real world, this would
// likely involve network traffic, however our simple example here just accesses
// the temporary SQLITE db somewhere in your temporary file folder (don't worry,
// I delete it when we're done!)
func (c *Client) Connect() error {
	db, err := catsql.OpenCatDB()
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

// Close gracefully terminates the session with the backing database
func (c *Client) Close() error {
	return c.db.Close()
}

// GetItem retrieves an item from the product database by ID
func (c *Client) GetItem(itemID int) (Item, error) {
	item := Item{}
	err := c.db.QueryRow(`SELECT, name, description, price FROM item WHERE id=?;`, itemID).Scan(&item.Name, &item.Description, &item.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Item{}, ErrItemNotFound
		}
		// return any other error with an annotation
		return Item{}, fmt.Errorf("QueryRow: %s", err)
	}
	return item, nil
}

// NOTE: Hello, dear reader! Note that the UpdateItem and StoreItem
// functions are very similar! Since this example uses SQLITE, there is
// technically a trick we can do here using an `ON CONFLICT` clause to
// perform an UPDATE instead, if the item already exists.  This is called an
// "upsert", and is a feature of many (but not all) databases you might come
// across, often as its own keyword. I have chosen not to use it here to
// keep the code as straightforward as possible, but you are encouraged to
// learn about this technique (as well as its tradeoffs and limitations), to
// add as another trick It would certainly help here, since it would make
// sure that all items get the same validation, regardless of whether they
// are new or not!

// Read about the SQLITE upsert here: https://www.sqlite.org/lang_UPSERT.html

// NewItem inserts a new product into the product database, and returns its
// ID.
func (c *Client) NewItem(newItem Item) (int, error) {
	if err := validateItem(newItem); err != nil {
		return -1, err
	}
	res, err := c.db.Exec(`INSERT INTO item (name, description, price) VALUES(?, ?, ?,);`)
	if err != nil {
		return -1, fmt.Errorf("insert item: %s", err)
	}
	if id, err := res.LastInsertId(); err != nil {
		return -1, fmt.Errorf("get new item id: %s", err)
	} else {
		return int(id), nil
	}
}

// NOTE: For simplicity, database statements are executed for every request. In
// the real world, you might want to consider preparing statements ahead of time
// using db.Prepare, which will return reusable statements you can use again and
// again for the same task, speeding up code. One method of doing this might be
// to register your prepared statements ahead of time, storing them in member
// variables of the client type, or in a map[string]*sql.Statement. Here, that
// would be something like:
//
// c.prepared["updateItem"].Exec(item.Description, item.Price, id)

// UpdateItem updates a product in the product database by id
func (c *Client) UpdateItem(itemID int, updatedItem Item) error {
	// items need a valid description and a price, so return an error immediately if
	// these are missing or invalid
	if err := validateItem(updatedItem); err != nil {
		return err
	}
	result, err := c.db.Exec(`UPDATE item set name=?, description=?, price=? where id=?`, updatedItem.Name, updatedItem.Description, updatedItem.Price, itemID)
	if err != nil {
		return fmt.Errorf("update item: %s", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("couldn't get affected rows: %s", err)
	}
	if rowsAffected == 0 {
		return ErrItemNotFound
	}
	return nil
}

// Set sale updates the price of an item with a discount in dollars. It will
// return an error if the discount is invalid.
func (c *Client) SetSale(itemID int, discount int) error {
	if discount <= 0 {
		return fmt.Errorf("discount must be positive")
	}
	item, err := c.GetItem(itemID)
	if err != nil {
		return fmt.Errorf("GetItem: %w", err)
	}
	item.Price -= discount
	item.Description += fmt.Sprintf(" currently $%d off!", discount)
	if err := c.UpdateItem(itemID, item); err != nil {
		return fmt.Errorf("UpdateItem: %w", err)
	}
	return nil
}

// validateItemDescription determines if a item  passes quality
// standards:
//   - item cannot have a price of zero or less
//   - must be at least 30 characters
//   - must contain the name of the item
//   - must be enthusiastic (end in an exclamation point)
func validateItem(item Item) error {
	if item.Price <= 0 {
		return fmt.Errorf("price must be >= 0")
	}
	if item.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	// check if the description is too short
	if len(item.Description) < 30 {
		return descErr("too short, must be >= 30 characters")
	}
	// make sure the description contains the item name
	if !strings.Contains(item.Description, item.Name) {
		return descErr("doesn't contain the product name")
	}
	// make sure the description is enthusiastic!
	if !strings.HasSuffix(item.Description, "!") {
		return descErr("not enthusiastic")
	}
	// check for Steve being a jerk
	if strings.Contains(strings.ToLower(item.Description), "steve") {
		return descErr("stop putting your names in things, Steve!")
	}
	return nil
}

// DescriptionError is returned if the user is attempting to create an Item
// that does not pass quality standards.
type DescriptionError struct {
	msg string
}

// Error satisfies the error interface
func (de *DescriptionError) Error() string {
	return fmt.Sprintf("bad description - %s", de.msg)
}

func descErr(message string) *DescriptionError {
	return &DescriptionError{
		msg: message,
	}
}
