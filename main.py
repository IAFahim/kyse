import keyboard
import threading


def key_remove(st, key):
    key.remove_abbreviation("e")


def key_set():
    keyboard.add_abbreviation("e", "fahimmanowarj5@gmail.com")
    keyboard.add_hotkey('ctrl+a', key_remove, args=('triggered',keyboard))
    keyboard.wait("F8")


if __name__ == "__main__":
    t1 = threading.Thread(key_set())
