using Microsoft.AspNetCore.Mvc;

namespace api_gateway.Controllers;

[ApiController, Route("[controller]")]
public class PingController : ControllerBase
{
    [HttpGet]
    public IActionResult Ping()
    {
        return Ok();
    }
}