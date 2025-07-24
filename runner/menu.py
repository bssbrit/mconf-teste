import json

def display_menu():
    """Display the main menu options"""
    print("\n=== Book Search Application ===")
    print("1. Search for a book")
    print("2. Exit")
    print("=" * 32)

def format_book_results(response_text):
    """Format the JSON response for better readability"""
    try:
        books = json.loads(response_text)
        if not books:
            return "No books found."
        
        result = f"\nFound {len(books)} book(s):\n"
        result += "=" * 50 + "\n"
        
        for i, book in enumerate(books, 1):
            result += f"{i}. Title: {book.get('title', 'N/A')}\n"
            authors = book.get('author_name', [])
            result += f"   Author(s): {', '.join(authors) if authors else 'N/A'}\n"
            result += f"   First Published: {book.get('first_publish_year', 'N/A')}\n"
            
            isbn = book.get('isbn', [])
            if isbn:
                result += f"   ISBN: {isbn[0]}\n"
            
            publishers = book.get('publisher', [])
            if publishers:
                result += f"   Publisher: {publishers[0]}\n"
            
            result += "-" * 30 + "\n"
        
        return result
    except json.JSONDecodeError:
        return f"Error parsing response: {response_text}"

def get_user_choice():
    """Get user menu choice"""
    try:
        choice = input("Enter your choice (1-2): ").strip()
        return int(choice)
    except ValueError:
        return 0

def get_book_name():
    """Get book name from user input"""
    book_name = input("Enter book name to search: ").strip()
    if not book_name:
        raise ValueError("Book name cannot be empty")
    return book_name

def ask_for_another_search():
    """Ask user if they want to perform another search"""
    while True:
        choice = input("\nWould you like to search for another book? (y/n): ").strip().lower()
        if choice in ['y', 'yes']:
            return True
        elif choice in ['n', 'no']:
            return False
        else:
            print("Please enter 'y' for yes or 'n' for no.")