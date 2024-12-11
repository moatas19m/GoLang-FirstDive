package models

import (
	"github.com/moatas19m/GoLang-FirstDive/pkg/config"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (b *Book) CreateBook() (*Book, error) {
	err := db.Create(b).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return b, nil
}

func GetAllBooks() ([]Book, error) {
	var books []Book
	err := db.Find(&books).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return books, nil
}

func GetBookById(id uint) (*Book, error) {
	var book Book
	err := db.First(&book, id).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &book, nil
}

func DeleteBookById(id uint) (Book, error) {
	var book Book
	err := db.Delete(&book, id).Error
	if err != nil {
		log.Fatal(err)
		return book, err
	}
	return book, nil
}

func UpdateBookById(id uint, updatedBook *Book) (Book, error) {
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		return book, err
	}

	book.Name = updatedBook.Name
	book.Author = updatedBook.Author
	book.Publication = updatedBook.Publication
	if err := db.Save(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

// While creating this file, I researched quite a bit on receivers (the (b *Book) between func and function name).
// Here are my findings:
//
// 1. **What is a Receiver?**
//    A receiver is a special parameter declared between the `func` keyword and the function name in Go.
//    It allows a function to be associated with a specific type, effectively making it a "method" of that type.
//
//    Example:
//        func (b *Book) CreateBook() (*Book, error) {
//            // 'b' is the receiver. It refers to the specific instance of 'Book' on which this method is called.
//        }
//
// 2. **Types of Receivers**
//    - **Pointer Receiver (`*Book`)**:
//        - The receiver is a pointer to the type (`*Book`).
//        - Any changes made to the fields of the receiver will persist outside the method.
//        - Avoids copying the struct, making it more efficient for larger types.
//
//        Example:
//            func (b *Book) UpdateName(newName string) {
//                b.Name = newName  // Updates the original Book instance
//            }
//
//    - **Value Receiver (`Book`)**:
//        - The receiver is passed by value.
//        - A copy of the type is made, so changes to the receiver inside the method do not affect the original instance.
//        - Suitable for small structs or methods that don’t modify the receiver.
//
//        Example:
//            func (b Book) PrintDetails() {
//                fmt.Println("Name:", b.Name)  // Prints details without modifying the original
//            }
//
// 3. **When to Use Receivers vs Regular Functions**
//    - **Use Receivers**:
//        - When the function logically belongs to the type (e.g., `CreateBook`, `UpdateBook`).
//        - When the function needs to operate on or modify the type's fields.
//
//    - **Use Regular Functions**:
//        - When the function doesn’t require access to the type’s instance (e.g., fetching all records from the database).
//        - For generic operations that are not specific to a particular instance.
//
// 4. **Counterarguments: Tightly Coupling Functionality to the Struct**
//    - **Why Keep Everything as a Method on `Book`?**
//        - It creates a more object-oriented design, where `Book` is the "owner" of all its operations.
//        - It makes your API intuitive for developers used to object-oriented paradigms, allowing calls like:
//            ```go
//            book.CreateBook()
//            book.UpdateBookById(&updatedBook)
//            ```
//        - This approach bundles all `Book`-related functionality into the `Book` struct itself.
//
//    - **Why Split Into Regular Functions and Receiver Methods?**
//        - Follows a cleaner **separation of concerns**.
//        - **Instance-Specific Logic** (e.g., `CreateBook`) belongs to the `Book` type, as it requires an instance.
//        - **Generic Operations** (e.g., `GetAllBooks`) belong at the database/repository level, as they don’t tie to a specific instance of `Book`.
//        - This approach aligns better with a functional programming mindset, avoiding unnecessary coupling.
//
// 5. **Why Use Pointer Receivers (`*Book`) Most of the Time?**
//    - To avoid copying the struct, which is particularly important for large structs.
//    - To allow methods to modify the fields of the struct directly.
//
// 6. **Practical Example**
//    struct Book {
//        ID          uint
//        Name        string
//        Author      string
//        Publication string
//    }
//
//    - Pointer Receiver Example:
//        func (b *Book) UpdateName(newName string) {
//            b.Name = newName // Updates the original struct
//        }
//
//    - Value Receiver Example:
//        func (b Book) PrintDetails() {
//            fmt.Println(b.Name) // Prints details without modifying the original struct
//        }
//
//    - Regular Function Example (No Receiver):
//        func GetAllBooks() ([]Book, error) {
//            var books []Book
//            err := db.Find(&books).Error
//            return books, err
//        }
//
// 7. **Best Practices**
//    - Use pointer receivers when the method modifies the struct or if the struct is large.
//    - Use value receivers for small, immutable structs or read-only methods.
//    - Avoid receivers when the function operates at the "global" level (e.g., fetching all records).
