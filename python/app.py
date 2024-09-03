import argparse
import os


APP_NAME = "CLI-Password Generator"
LOWER_CHARS = "abcdefghijklmnopqrstuvwxyz"
UPPER_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
NUMBER_CHARS = "0123456789"
SYMBOL_CHARS = "!@#$%^&*_-+="
ALPHA_CHARS = upperChars + lowerChars


def create_parser():
    parser = argparse.ArgumentParser(description=APP_NAME)
