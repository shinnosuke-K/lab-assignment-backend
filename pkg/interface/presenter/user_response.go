package presenter

import (
	"github.com/shinnosuke-K/lab-assignment-backend/pkg/domain/model"
	"github.com/shinnosuke-K/lab-assignment-backend/pkg/interface/controller"
)

func UserOKResponse(user *model.User, professors []*model.Professor, answers []*model.Answer) controller.UserViewResponse {

	var (
		labNames               []string
		professorsByLabNameMap = map[string][]*model.Professor{}
		pointByProfIDMap       = map[string]int{}
	)

	for _, prof := range professors {
		professorsByLabNameMap[prof.LabName] = append(professorsByLabNameMap[prof.LabName], prof)
	}

	for key, _ := range professorsByLabNameMap {
		labNames = append(labNames, key)
	}

	for _, ans := range answers {
		pointByProfIDMap[ans.ProfID] = ans.Point
	}

	return controller.UserViewResponse{
		UserID:   user.ID,
		Graduate: user.Graduate,
		Labs:     newLabs(labNames, professorsByLabNameMap, pointByProfIDMap),
	}
}

func newLabs(labNames []string, professorsByLabNameMap map[string][]*model.Professor, pointByProfIDMap map[string]int) []*controller.Lab {
	res := make([]*controller.Lab, len(labNames))
	for i, l := range labNames {
		res[i] = newLab(l, professorsByLabNameMap[l], pointByProfIDMap)
	}
}

func newLab(labName string, professors []*model.Professor, pointByProfIDMap map[string]int) *controller.Lab {
	return &controller.Lab{
		Name:       labName,
		Professors: newProfessors(professors, pointByProfIDMap),
	}
}

func newProfessors(professors []*model.Professor, pointByProfIDMap map[string]int) []*controller.Professor {
	res := make([]*controller.Professor, len(professors))
	for i, prof := range professors {
		res[i] = newProfessor(prof.ProfName, pointByProfIDMap[prof.ID])
	}
}

func newProfessor(name string, point int) *controller.Professor {
	return &controller.Professor{
		Name:  name,
		Point: point,
	}
}
