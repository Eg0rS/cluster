package query

const GetRequestSql = `
select PersonnelRequest.request_title,
       request_description,
       PersonnelRequest.test_id,
       users.user_name,
       users.surname
from PersonnelRequest
join Users on $1 = Users.id
where PersonnelRequest.user_id = $1;
`

const GetAllTestIdsByUserIdSql = `
select test_id, p.test_type
from PersonnelRequest
         join PersonnelTest p on p.id = PersonnelRequest.test_id
where user_id = $1;
`
const GetTestInfoByTestId = `
SELECT test_title, test_description, test_type from PersonnelTest where id = $1;
`

const GetQuestionsByTestId = `
SELECT id, question_text  FROM TestQuestion WHERE test_id = $1;
`

const GetAnswersByQuestionId = `
SELECT answer, is_right FROM QuestionAnswer where question_id = $1;
`

const GetOrganizationsSql = `
	SELECT org_name, address from Organizations;
`
