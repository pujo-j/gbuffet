import { Component, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import {animate, state, style, transition, trigger} from '@angular/animations';
import { AddDialogComponent } from './add-dialog/add-dialog.component';
import { Router } from '@angular/router';
export interface UserData {
  id: string;
  project_id: string;
  project_name: string;
  creation_date: string;
  status: string;
  requester_email: string
}


/**
 * @title Data table with sorting, pagination, and filtering.
 */
@Component({
  selector: 'app-project-request',
  templateUrl: './project-request.component.html',
  styleUrls: ['./project-request.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({height: '0px', minHeight: '0', display: 'none'})),
      state('expanded', style({height: '*'})),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
export class ProjectRequestComponent implements OnInit {
  private addDialog: MatDialogRef<AddDialogComponent>;
  dialogStatus = 'inactive';
  isAdminMode: boolean = false;
  userTable: string[] = ['id', 'project_id', 'project_name', 'creation_date', 'status'];
  adminTable: string[] = ['id', 'requester_email', 'project_id', 'project_name', 'creation_date', 'status']
  displayedColumns: string[] = this.isAdminMode != true? this.userTable: this.adminTable;
  dataSource: MatTableDataSource<UserData>;
  expandedElement: UserData | null;
  

  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(MatSort) sort: MatSort;

  constructor(public dialog: MatDialog, private router: Router) {
    // get project http request here
    const users: UserData [] = [
      {
        id: "1",
        project_id: "gbuffet_test",
        project_name: "Gcp Buffet",
        creation_date: "2016-1-17",
        status: "NEW",
        requester_email: "andresse@gmail.fr"
      },
      {
        id: "3",
        project_id: "Aliast_test",
        project_name: "Allea",
        creation_date: "2016-3-17",
        status: "GRANTED",
        requester_email: "andresse@gmail.fr"
      },
      {
        id: "2",
        project_id: "Aliast_test",
        project_name: "Bllea",
        creation_date: "2016-2-17",
        status: "GRANTED",
        requester_email: "andresse@gmail.fr"
      }
  ]

    // Assign the data to the data source for the table to render
    this.dataSource = new MatTableDataSource(users);
  }

  ngOnInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  applyFilter(filterValue: string) {
    this.dataSource.filter = filterValue.trim().toLowerCase();

    if (this.dataSource.paginator) {
      this.dataSource.paginator.firstPage();
    }
  }

  
  // Open dialog
  showDialog() {
    this.dialogStatus = 'active';
    this.addDialog = this.dialog.open(AddDialogComponent, {
      width: '450px',
      data: {}
    });

    this.addDialog.afterClosed().subscribe(project_request => {
      this.dialogStatus = 'inactive';
      console.log('request: ', project_request);
      if (project_request) {
        //this.add(project_request);
      }
    });
  }

  hideDialog() {
    this.dialogStatus = 'inactive';
    this.addDialog.close();
  }

}
