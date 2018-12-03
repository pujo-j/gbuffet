package com.sfeir.open.gbuffetcontroltower

import com.vladmihalcea.hibernate.type.array.StringArrayType
import org.hibernate.annotations.Type
import org.hibernate.annotations.TypeDef
import org.hibernate.annotations.TypeDefs
import java.util.*
import javax.persistence.*

enum class RequestStatus {
    NEW, GRANTED, REJECTED
}

@TypeDefs(TypeDef(name = "string-array", typeClass = StringArrayType::class))
@Entity
data class ProjectRequest(

        @Id
        val id: String,

        @Column(nullable = false)
        val requester_email: String,

        @Column(nullable = false)
        val requester_group: String? = null,

        @Column(nullable = false)
        val expected_lifetime: Int = 1,

        @Column(nullable = false)
        val request_status: RequestStatus = RequestStatus.NEW,

        @Column(nullable = false)
        val creation: Date = Date(),

        @Column(nullable = false)
        val requester_comment: String = "",

        @Column(nullable = false)
        val admin_comment: String = "",

        @Type(type = "string-array")
        @Column(columnDefinition = "text[]")
        val tags: Set<String> = setOf(),

        @OneToOne
        val project: Project? = null,

        @OneToMany
        val allocations: List<Allocation> = listOf()
)