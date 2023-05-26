namespace curator_service.Models
{
    public class TraineeRequest
    {
        public bool IsRecommended { get; set; }
        public int Age { get; set; }
        public string InvitationSource { get; set; }
        public string TraineeRole { get; set; }
        public string City { get; set; }
        public string University { get; set; }
        public string Education { get; set; }
        public DateTime StartExperience { get; set; }
        public DateTime EndExperience { get; set; }
    }
}
