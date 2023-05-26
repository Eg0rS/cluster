using curator_service.Models;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace curator_service.Controllers
{
    [Route("curator-service")]
    // [Authorize]
    public class CuratorController : Controller
    {
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
