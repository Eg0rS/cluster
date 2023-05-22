using Microsoft.AspNetCore.Mvc;

namespace user_service.Controllers;

[ApiController, Route("[controller]")]
public class PingController : ControllerBase
{
    [HttpGet]
    public IActionResult Ping()
    {
        return Ok();
    }
}