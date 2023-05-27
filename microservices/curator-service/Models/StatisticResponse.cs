namespace curator_service.Models
{
    enum Education
    {

    }

    public class StatisticResponse
    {
        public int TraineeRequestsCount { get; set; }
        public List<(int, int)> AverageAge { get; set; }
        public List<(string, int)> Programs { get; set; }
        public List<(string, int)> Universities {get; set;}
        public List<(string, int)> Educations { get; set; }
        // статистика по опыту работы
        // направления стажировки
        // каналы привлечения
        public int CardRequestsCount { get; set; }
        public int AuthorityRequestsCount { get; set;}
        // статистика по образованию
    }
}
