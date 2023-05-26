package personnel

import "personnel_service/model"

func MapRequestRadioTestModelToServiceModel(testModel RadioTest) model.RadioTest {
	var questions []model.Question

	for i := range testModel.Questions {
		var answers []model.Answer

		for l := range testModel.Questions[i].Answers {
			answers = append(answers, model.Answer{Text: testModel.Questions[i].Answers[l].Text, IsRight: testModel.Questions[i].Answers[l].IsRight})
		}

		questions = append(questions, model.Question{Title: testModel.Questions[i].Title, Answers: answers})
	}

	return model.RadioTest{
		Title:       testModel.Title,
		Description: testModel.Description,
		Questions:   questions,
	}
}

func MapServiceRadioTestModelToRequestModel(testModel model.RadioTest) RadioTest {
	var questions []Question

	for i := range testModel.Questions {
		var answers []Answer

		for l := range testModel.Questions[i].Answers {
			answers = append(answers, Answer{Text: testModel.Questions[i].Answers[l].Text, IsRight: testModel.Questions[i].Answers[l].IsRight})
		}

		questions = append(questions, Question{Title: testModel.Questions[i].Title, Answers: answers})
	}

	return RadioTest{
		Title:       testModel.Title,
		Description: testModel.Description,
		Questions:   questions,
	}
}

func MapRequestRequestModelToServiceRequestModel(modelReq Request) model.Request {
	return model.Request{
		Title:          modelReq.Title,
		Description:    modelReq.Description,
		TestId:         modelReq.TestId,
		UserId:         modelReq.UserId,
		OrganizationId: modelReq.Organization,
	}
}

func MapRequestTextTestModelToServiceModel(modelReq TextTest) model.CreateTextTest {
	return model.CreateTextTest{
		Title:       modelReq.Title,
		Description: modelReq.Description,
		File:        modelReq.File,
		Header:      modelReq.Header,
	}
}
