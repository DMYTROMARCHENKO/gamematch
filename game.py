import time
import random

def welcome_message():
    print("Ласкаво просимо до математичної гри!")

def countdown():
    print("Відлік до початку гри:")
    for i in range(5, 0, -1):
        print(i)
        time.sleep(1)

def generate_question():
    num1 = random.randint(1, 10)
    num2 = random.randint(1, 10)
    return num1, num2

def validate_answer(answer, num1, num2):
    correct_answer = num1 + num2
    return answer == correct_answer

def main():
    welcome_message()
    countdown()

    score = 0
    start_time = time.time()

    while score < 50:
        num1, num2 = generate_question()
        print(f"Що {num1} + {num2}?")
        answer = int(input("Ваша відповідь: "))
        
        if validate_answer(answer, num1, num2):
            score += 5
            print(f"Правильно! Ваш поточний рахунок {score}. Тобі потрібно {50 - score} очок, щоб закінчити гру.")
        else:
            print("Невірно! Спробуйте знову.")

    end_time = time.time()
    time_taken = end_time - start_time
    print(f"Щиро вітаю! Ви закінчили гру за {time_taken:.2f} секунд.")

if __name__ == "__main__":
    main()
