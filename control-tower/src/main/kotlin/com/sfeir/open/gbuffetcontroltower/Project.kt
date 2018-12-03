package com.sfeir.open.gbuffetcontroltower

import org.hibernate.annotations.Type
import java.util.*
import javax.persistence.*

@Entity
data class Project(
        @Id
        val id: String,

        @Column(nullable = false)
        val name: String,

        @Column(nullable = false)
        val project_number: String,

        @OneToOne(mappedBy = "project")
        val request: ProjectRequest,

        @Column(nullable = false)
        val creation: Date = Date(),

        @Column(nullable = false)
        val folder: String,

        @Type(type = "string-array")
        @Column(columnDefinition = "text[]")
        val tags: Set<String> = setOf(),

        @OneToMany
        val allocations: Set<Allocation> = setOf()
)