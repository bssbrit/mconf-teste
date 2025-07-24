import requests
import sys
from menu import display_menu, format_book_results, get_user_choice, get_book_name, ask_for_another_search

def search_book(book_name):
   
    print(f"Searching for: {book_name}")
    
    params = {'book_name': book_name}
    try:
        response = requests.get('http://localhost:3000/', params=params)
        
        if response.status_code == 200:
            formatted_response = format_book_results(response.text)
            print(formatted_response)
        else:
            print(f"Error: {response.status_code} - {response.text}")
    except requests.RequestException as e:
        print(f"Connection error: {e}")

def main():
   
    if len(sys.argv) > 1:
        book_name = " ".join(sys.argv[1:])
        search_book(book_name)
        
    
        while ask_for_another_search():
            try:
                book_name = get_book_name()
                search_book(book_name)
            except ValueError as e:
                print(f"Error: {e}")
        
        print("Goodbye!")
        return
    
    while True:
        display_menu()
        choice = get_user_choice()
        
        if choice == 1:
            try:
                book_name = get_book_name()
                search_book(book_name)
                
                
                if not ask_for_another_search():
                    print("Goodbye!")
                    break
            except ValueError as e:
                print(f"Error: {e}")
        elif choice == 2:
            print("Goodbye!")
            break
        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main()