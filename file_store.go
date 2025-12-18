package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

/*
File format (users.txt):
id,name,email
*/

func readUsersFromFile() ([]User, error) {
	file, err := os.Open("users.txt")
	if err != nil {
		// If file does not exist, return empty slice (not error)
		return []User{}, nil
	}
	defer file.Close()

	var users []User
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) < 3 {
			continue
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		users = append(users, User{
			ID:    id,
			Name:  parts[1],
			Email: parts[2],
		})
	}

	// Always return non-nil slice
	if users == nil {
		users = []User{}
	}

	return users, nil
}

func writeUsersToFile(users []User) error {
	file, err := os.Create("users.txt") // overwrite file
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, user := range users {
		line := strconv.Itoa(user.ID) + "," + user.Name + "," + user.Email + "\n"
		if _, err := writer.WriteString(line); err != nil {
			return err
		}
	}

	return writer.Flush()
}
