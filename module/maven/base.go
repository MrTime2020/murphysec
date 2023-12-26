package maven

import (
	"context"
	"github.com/murphysecurity/murphysec/model"
	"github.com/murphysecurity/murphysec/utils"
	"path/filepath"
)

type Inspector struct{}

func (Inspector) SupportFeature(feature model.InspectorFeature) bool {
	return false
}

func (Inspector) String() string {
	return "Maven"
}

func (Inspector) CheckDir(dir string) bool {
	return utils.IsFile(filepath.Join(dir, "pom.xml"))
}

func (Inspector) InspectProject(ctx context.Context) error {
	task := model.UseInspectionTask(ctx)
	modules, e := ScanMavenProject(ctx, task)
	if e != nil {
		return e
	}
	for _, it := range modules {
		task.AddModule(it)
	}
	return nil
}
