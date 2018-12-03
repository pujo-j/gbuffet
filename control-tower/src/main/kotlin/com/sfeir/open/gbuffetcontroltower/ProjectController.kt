package com.sfeir.open.gbuffetcontroltower

import org.springframework.boot.autoconfigure.EnableAutoConfiguration
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RequestMethod
import org.springframework.web.bind.annotation.RestController
import java.util.*

@RestController
@EnableAutoConfiguration
class ProjectController {
    @RequestMapping("/project/{id}", method = [RequestMethod.GET])
    fun project(): Project {
        val pr = ProjectRequest(
                id = "test",
                requester_email = "test@example.com",
                requester_group = null,
                expected_lifetime = 1,
                request_status = RequestStatus.NEW,
                creation = Date(),
                requester_comment = "No comment",
                admin_comment = "No comment either",
                tags = setOf("toto", "tata", "titi")
        )
        return Project(
                id = "test",
                name = "test",
                project_number = "0000000"
        )
    }
}