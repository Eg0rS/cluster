package pesrsonnel_repo

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"personnel_service/database/pesrsonnel_repo/query"
	"personnel_service/model"
)

type PersonnelRepository struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

func NewPersonnelRepository(logger *zap.SugaredLogger, db *sqlx.DB) PersonnelRepository {
	return PersonnelRepository{
		logger: logger,
		db:     db,
	}
}

func (r PersonnelRepository) InsertRadioTest(ctx context.Context, testModel model.RadioTest) (int, error) {
	var (
		psql       = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
		testId     int
		questionId int
	)

	builder := psql.Insert("PersonnelTest").
		Columns("test_title", "test_description", "test_type").
		Values(testModel.Title, testModel.Description, "radio_test").
		Suffix("RETURNING \"id\"").
		RunWith(r.db)
	err := builder.QueryRowContext(ctx).Scan(&testId)
	if err != nil {
		return 0, err
	}

	for i := range testModel.Questions {
		builder := psql.Insert("TestQuestion").
			Columns("question_text", "test_id").
			Values(testModel.Questions[i].Title, testId).
			Suffix("RETURNING \"id\"").
			RunWith(r.db)
		err := builder.QueryRowContext(ctx).Scan(&questionId)
		if err != nil {
			return 0, err
		}

		for l := range testModel.Questions[i].Answers {
			builder := psql.Insert("QuestionAnswer").
				Columns("answer", "is_right", "question_id").
				Values(testModel.Questions[i].Answers[l].Text, testModel.Questions[i].Answers[l].IsRight, questionId).
				RunWith(r.db)
			sqlQuery, args, err := builder.ToSql()
			if err != nil {
				return 0, err
			}

			_, err = r.db.ExecContext(ctx, sqlQuery, args...)
			if err != nil {
				return 0, err
			}
		}
	}

	return testId, nil
}

func (r PersonnelRepository) InsertTextTest(ctx context.Context, testModel model.CreateTextTest, filePath string) (int, error) {
	var (
		psql   = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
		testId int
	)

	builder := psql.Insert("PersonnelTest").
		Columns("test_title", "test_description", "test_type").
		Values(testModel.Title, testModel.Description, "text_test").
		Suffix("RETURNING \"id\"").
		RunWith(r.db)
	err := builder.QueryRowContext(ctx).Scan(&testId)
	if err != nil {
		return 0, err
	}

	builder = psql.Insert("TestFile").
		Columns("filename", "test_id").
		Values(filePath, testId).
		RunWith(r.db)

	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	_, err = r.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return 0, err
	}

	return testId, nil
}

func (r PersonnelRepository) InsertRequest(ctx context.Context, reqModel model.Request) error {
	var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	builder := psql.Insert("PersonnelRequest").
		Columns("request_title", "request_description", "test_id", "user_id", "organization_id").
		Values(reqModel.Title, reqModel.Description, reqModel.TestId, reqModel.UserId, reqModel.OrganizationId).
		RunWith(r.db)
	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r PersonnelRepository) GetAllRequests(ctx context.Context, userId string) ([]Request, error) {
	var requests []Request

	rows, err := r.db.QueryContext(ctx, query.GetRequestSql, userId)
	if err != nil {
		return []Request{}, err
	}

	for rows.Next() {
		var (
			name        string
			surname     string
			title       string
			description string
			testId      int
		)

		err = rows.Scan(&title, &description, &testId, &name, &surname)
		if err != nil {
			return []Request{}, err
		}

		requests = append(requests, Request{
			Title:       title,
			Description: description,
			Name:        name,
			Surname:     surname,
			TestId:      testId,
		})
	}

	return requests, nil
}

func (r PersonnelRepository) GetAllTests(ctx context.Context, userId int) (model.AllTests, error) {
	var (
	//res model.AllTests
	)

	rows, err := r.db.QueryContext(ctx, query.GetAllTestIdsByUserIdSql, userId)
	if err != nil {
		return model.AllTests{}, err
	}

	for rows.Next() {
		var (
			testType string
			testId   int
		)

		err := rows.Scan(&testId, &testType)
		if err != nil {
			return model.AllTests{}, err
		}

		if testType == "text_test" {

		}

		if testType == "radio_test" {

		}
	}

	return model.AllTests{}, nil
}

func (r PersonnelRepository) GetTestById(ctx context.Context, testId string) (model.RadioTest, error) {
	var (
		testType        string
		testTitle       string
		testDescription string
		questionId      int
		questionText    string
		data            model.RadioTest
	)

	rows, err := r.db.QueryContext(ctx, query.GetTestInfoByTestId, testId)
	if err != nil {
		return model.RadioTest{}, err
	}

	if !rows.Next() {
		return model.RadioTest{}, errors.New("Test isn't exists")
	}

	for rows.Next() {
		err = rows.Scan(&testTitle, &testDescription, &testType)
		if err != nil {
			return model.RadioTest{}, err
		}
	}

	data.Title = testTitle
	data.Description = testDescription

	if testType == "radio_test" {
		rows, err = r.db.QueryContext(ctx, query.GetQuestionsByTestId, testId)
		if err != nil {
			return model.RadioTest{}, err
		}

		var questions []model.Question

		for rows.Next() {
			var question model.Question

			err := rows.Scan(&questionId, &questionText)
			if err != nil {
				return model.RadioTest{}, err
			}

			question.Title = questionText

			answerRows, err := r.db.QueryContext(ctx, query.GetAnswersByQuestionId, questionId)
			if err != nil {
				return model.RadioTest{}, err
			}

			var answers []model.Answer

			for answerRows.Next() {
				var answerTitle string
				var isRight bool

				err := answerRows.Scan(&answerTitle, &isRight)
				if err != nil {
					return model.RadioTest{}, err
				}

				var answer = model.Answer{Text: answerTitle, IsRight: isRight}
				answers = append(answers, answer)
			}

			question.Answers = answers
			questions = append(questions, question)
		}

		data.Questions = questions
		data.TestType = "radio_test"

		return data, nil
	}

	if testType == "text_test" {

	}

	return model.RadioTest{}, nil
}

func (r PersonnelRepository) InsertOrganization(ctx context.Context, organizationModel model.AddOrganizationModel) (int, error) {
	var (
		psql           = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
		organizationId int
	)

	builder := psql.Insert("Organizations").
		Columns("org_name", "address", "first_coordinates", "second_coordinates").
		Values(organizationModel.Name, organizationModel.Address, organizationModel.FirstCoordinate, organizationModel.SecondCoordinate).
		Suffix("RETURNING \"id\"").
		RunWith(r.db)
	err := builder.QueryRowContext(ctx).Scan(&organizationId)
	if err != nil {
		return 0, err
	}

	return organizationId, nil
}

func (r PersonnelRepository) SelectOrganizations(ctx context.Context) (model.GetOrganizationsModel, error) {
	var (
		organisations []model.OrganizationInfo
	)

	rows, err := r.db.QueryContext(ctx, query.GetOrganizationsSql)
	if err != nil {
		return model.GetOrganizationsModel{}, err
	}

	for rows.Next() {
		var (
			name    string
			address string
			x       float64
			y       float64
		)

		err := rows.Scan(&name, &address, &x, &y)
		if err != nil {
			return model.GetOrganizationsModel{}, err
		}

		org := model.OrganizationInfo{Address: address, Name: name, X: x, Y: y}
		organisations = append(organisations, org)
	}

	return model.GetOrganizationsModel{Organizations: organisations}, nil
}
