using Microsoft.AspNetCore.Mvc;

namespace auth.Controllers;

[ApiController, Route("[controller]")]
public class PingController : ControllerBase
{
    [HttpGet]
    public IActionResult Ping()
    {
        return Ok();
    }
}