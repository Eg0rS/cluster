namespace curator_service.Models
{
    public class Teacher
    {
        public string Surname { get; set; }
        public string Name { get; set; }
        public string Patronymic { get; set; }
        public bool IsCompletedSchool { get; set; }
        public bool HasSertificate { get; set; }
        public int Rating { get; set; }
    }
}
