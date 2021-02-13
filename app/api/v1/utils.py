import hashlib
from random import randint
from faker import Faker

LETTER = 'abcdefghijklmnopqrstuvwxyz'
CASE_LETTER = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'

Faker.seed(0)
faker = Faker()


def hash_generator(pin_len, salt_len=10, strong=False):
    """
    Generate PIN code, salt and SHA-1 hash pin+salt
    :param pin_len Length of PIN code. Default=6
    :param salt_len Lenght of salt. Default=10
    :param strong True if need letter in pin. Default=False
    """
    # Create number PIN code
    pin = []
    if strong:
        for _ in range(pin_len-2):
            pin.append(str(randint(0, 9)))
        pin = symbol_add(pin)
    else:
        for _ in range(pin_len):
            pin.append(str(randint(0, 9)))


    # Generate random salt
    salt = faker.hexify(text=f'^'*salt_len, upper=False)

    # If need letters add


    # Create password
    password = ''.join(pin) + salt

    # Getting hash of password
    h = hashlib.sha1(password.encode("utf-8"))

    return {
        "pin": ''.join(pin),
        "salt": salt,
        "hash": h.hexdigest().upper()
    }


def symbol_add(pin):
    for letters in [LETTER, CASE_LETTER]:
        letter = letters[randint(0, len(letters)-1)]
        letter_position = randint(0, len(pin))
        pin.insert(letter_position, letter)

    return pin
