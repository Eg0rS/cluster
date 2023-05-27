using curator_service.Models;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace curator_service.Controllers
{
    [Route("curator-service")]
    // [Authorize]
    public class CuratorController : Controller
    {
        /// <summary>
        /// Выводит список поданых заявок с фильтром.
        /// </summary>
        /// <param name="filter">Фильтрация заявок. Возможные значения: recommended, non-recommended, не указано (без фильтра)</param>
        /// <response code="200">Успех. Выводится список заявок.</response>
        /// <response code="500">Внутренняя ошибка. Не были получены заявки от стороннего МКС-а.</response>
        [Route("trainee-requests")]
        [HttpGet]
        public JsonResult GetTraineeRequests([FromQuery] string filter)
        {
            var requests = new List<TraineeRequest> { };
            // здесь обращение к мкс стажеров и получение всех заявок без фильтра + валидация
            if (requests == null)
            {
                return new JsonResult(StatusCode(500));
            }
            return filter == "recommended"
                ? new JsonResult(requests.Where(x => x.IsRecommended))
                : filter == "non-recommended"
                    ? new JsonResult(requests.Where(x => !x.IsRecommended))
                    : new JsonResult(requests);
        }


        /// <summary>
        /// Выводит список поданых заявок на стажеров от кадров.
        /// </summary>
        /// <response code="200">Успех. Выводится список заявок.</response>
        /// <response code="500">Внутренняя ошибка. Не были получены заявки от стороннего МКС-а.</response>
        [Route("card-requests")]
        [HttpGet]
        public JsonResult GetCadrRequests()
        {
            var requests = new List<CadrRequest> { };
            // обращение к мкс
            if (requests == null)
            {
                return new JsonResult(StatusCode(500));
            }

            return new JsonResult(requests);
        }

        /// <summary>
        /// Редактирует заявки на стажеров от кадров. 
        /// </summary>
        /// <param name="request">Обновленная заявка на стажера.</param>
        /// <response code="200">Успех. Заявка была успешно изменена.</response>
        /// <response code="500">Внутренняя ошибка. Не была получена информация от стороннего МКС-а.</response>
        [Route("card-requests")]
        [HttpPut]
        public JsonResult UpdateCadrRequests([FromBody] CadrRequest request)
        {
            var requests = new List<CadrRequest> { };
            // обращение к мкс
            if (requests == null)
            {
                return new JsonResult(StatusCode(500));
            }

            var cadrRequest = requests.FirstOrDefault(x => x.Surname == request.Surname && x.Name == request.Name && x.Patronymic == request.Patronymic);
            if (cadrRequest == null)
            {
                return new JsonResult(StatusCode(500));
            }

            // изменение каких надо полей кадрреквеста как надо
            // обращение к мкс

            return new JsonResult(StatusCode(200));
        }

        /// <summary>
        /// Выводит список наставников, которые прошли обучение в школе наставников, а так же информацию о них.
        /// </summary>
        /// <response code="200">Успех.</response>
        /// <response code="500">Внутренняя ошибка. Не была получена информация от стороннего МКС-а.</response>
        [Route("teachers")]
        [HttpGet]
        public JsonResult GetTeachersData()
        {
            var teachers = new List<Teacher> { };
            // обращение к мкс

            if(teachers == null)
            {
                return new JsonResult(StatusCode(500));
            }

            return new JsonResult(teachers.Where(
                    x => x.IsCompletedSchool
                ));
        }

        /// <summary>
        /// Выводит статистику по заявкам от кандидатов на стажеров и от стажеров.
        /// </summary>
        /// <response code="200">Успех.</response>
        /// <response code="500">Внутренняя ошибка. Не была получена информация от стороннего МКС-а.</response>

        [Route("statistics")]
        [HttpGet]
        public JsonResult GetStatistics()
        {
            var requests = new List<TraineeRequest> { };
            var trainees = new List<Trainee> { };
            var cardRequests = new List<CadrRequest> { };

            // обращение к мкс и получение данных в листы, валидация

            // заявки НА стажеров
            var requestsCount = requests.Count; // количество всех откликов
            var ages = new List<(int, int)> { }; // количество всех возрастов
            var traineePrograms = new List<(string, int)> { }; // направления стажировок
            var invitations = new List<(string, int)> { }; // от кого участники узнали о программах стажировок
            var cities = new List<(string, int)> { }; // статистика городов
            var universities = new List<(string, int)> { }; // статистика ВУЗов
            var educations = new List<(string, int)> { }; // статистика образования
            var experience = new List<(string, int)> { ("AboveYear", 0), ("BelowYear", 0)}; // статистика по опыту работы
            for (int i = 0; i < requestsCount; i++)
            {
                ages = AddToTuple(ages, requests[i].Age);
                invitations = AddToTuple(invitations, requests[i].InvitationSource);
                traineePrograms = AddToTuple(traineePrograms, requests[i].TraineeRole);
                cities = AddToTuple(cities, requests[i].City);
                universities = AddToTuple(universities, requests[i].University);
                educations = AddToTuple(educations, requests[i].Education);
                var diff = (requests[i].EndExperience - requests[i].StartExperience).Days;
                experience = AddExperienceData(experience, diff);
            }

            // заявки стажеров
            var traineesCount = trainees.Count; // количество заявок на стажеров 
            var cadrRequestsCount = cardRequests.Count; // количество заявок от органов власти и учреждений
            var traineesEducation = new List<(string, int)> { }; // сбор информации об образовании стажеров
            for(int i = 0; i < trainees.Count; i++)
            {
                traineesEducation = AddToTuple(traineesEducation, trainees[i].Education);
            }

            return new JsonResult(new StatisticResponse
            {
                TraineeRequestsCount = requestsCount,
                AverageAge = ages,
                Programs = traineePrograms,
                Universities = universities,
                Educations = educations,

            });
        }

        private List<(T, int)> AddToTuple<T>(List<(T, int)> tuples, T data)
        {
            for(int i = 0; i<tuples.Count; i++)
            {
                if (tuples[i].Item1.Equals(data))
                {
                    var count = tuples[i].Item2 + 1;
                    tuples[i] = (data, count);
                    return tuples;
                }
            }

            tuples.Add((data, 1));
            return tuples;
        }

        private List<(string, int)> AddExperienceData(List<(string, int)> tuples, int data)
        {
            if(data <= 365)
            {
                return new List<(string, int)> { (tuples[0].Item1, tuples[0].Item2), (tuples[1].Item1, tuples[1].Item2 + 1)};
            }

            return new List<(string, int)> { (tuples[0].Item1, tuples[0].Item2 + 1), (tuples[1].Item1, tuples[1].Item2) };
        }
    }
}
