package utilities

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Jon1105/pmag/conf"
)

// String slice functions

func GetLanguage(str string, languages []conf.Language) (conf.Language, error) {
	var langString string = strings.ToLower(str)

	for _, v := range languages {
		if Contains(v.Acros, langString) {
			return v, nil
		}
	}

	return conf.Language{}, fmt.Errorf("invalid language name %q", str)
}

func InferLanguage(args []string, config *conf.Config) (string, conf.Language, error) {
	var name = args[0]

	for _, lang := range config.Languages {
		var projects, err = GetProjects(lang.Path)
		if err != nil {
			continue
		}
		for _, v := range projects {
			if v.Name() == name {
				return filepath.Join(lang.Path, v.Name()), lang, nil
			}
		}
	}

	return "", conf.Language{}, fmt.Errorf("could not infer project")
}

func Open(projectPath, editorPath string, disableExtensions bool) error {
	// if strings.HasSuffix(editorPath, "/bin/code") {
	// 	return fmt.Errorf("Project at path %q cannot be opened Visual Studio Code, due to inefficient energy consumption.", projectPath)
	// } else
	if strings.HasSuffix(editorPath, "Visual Studio Code.app") {
		fmt.Print("opening ", projectPath, " with ", editorPath)
		return RunCommand("", "open", editorPath)
	} else if disableExtensions {
		return RunCommand("", editorPath, projectPath, "--disable-extensions")
	} else {
		return RunCommand("", editorPath, projectPath)
	}
}

func GetEditorPath(langPath, defaultPath string) (string, error) {
	if langPath != "" {
		return langPath, nil
	} else if defaultPath != "" {
		return defaultPath, nil
	} else {
		return "", fmt.Errorf("DefaultEditorPath must not be empty")
	}
}
