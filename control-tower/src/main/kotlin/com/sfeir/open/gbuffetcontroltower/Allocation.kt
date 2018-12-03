package com.sfeir.open.gbuffetcontroltower

import javax.persistence.*

@Entity
data class Allocation(
        @Id
        @GeneratedValue(strategy = GenerationType.AUTO)
        val id: Long,

        @Column(nullable = false)
        val type: String,

        @Column(nullable = false)
        val unit: String,

        @Column(nullable = true)
        val region: String?,

        @Column(nullable = true)
        val zone: String?,

        @Column(nullable = false)
        val allocation: Long
)