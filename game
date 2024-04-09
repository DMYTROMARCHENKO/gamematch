import time
import random

def welcome_message():
    print("Ласкаво просимо до математичної гри!")

def countdown():
    def count_and_sleep(i):
        print(i)
        time.sleep(1)
    
    print("Відлік до початку гри:")
    list(map(count_and_sleep, range(5, 0, -1)))

def generate_question():
    return random.randint(1, 10), random.randint(1, 10)

def validate_answer(answer, num1, num2):
    return answer == num1 + num2

def main():
    welcome_message()
    countdown()

    score, start_time = 0, time.time()

    while score < 50:
        num1, num2 = generate_question()
        print(f"Що {num1} + {num2}?")
        answer = int(input("Ваша відповідь: "))
        
        if validate_answer(answer, num1, num2):
            score += 5
            print(f"Правильно! Ваш поточний рахунок {score}. Тобі потрібно {50 - score} очок, щоб закінчити гру.")
        else:
            print("Невірно! Спробуйте знову.")

    print(f"Ви закінчили гру за {time.time() - start_time:.2f} секунд.")

if __name__ == "__main__":
    main()
