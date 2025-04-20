package util

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReadTextFile 读取文本文件
func ReadTextFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read text file: %w", err)
	}
	return string(data), nil
}

// WriteTextFile 写入文本文件
func WriteTextFile(filePath string, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write text file: %w", err)
	}
	return nil
}

// ReadJSONFile 读取 JSON 文件
func ReadJSONFile(filePath string, target interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %w", err)
	}
	err = json.Unmarshal(data, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}
	return nil
}

// WriteJSONFile 写入 JSON 文件
func WriteJSONFile(filePath string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}
	return nil
}
