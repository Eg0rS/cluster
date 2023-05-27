namespace curator_service.Models
{
    public class CadrRequest
    {
        public string Surname { get; set; }
        public string Name { get; set; }
        public string Patronymic { get; set; }
        public bool IsConfirmed { get; set; }
    }
}
