import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AdminComponent } from './admin/admin.component';
import { ProjectRequestComponent } from './project-request/project-request.component';
import { AddDialogComponent } from './project-request/add-dialog/add-dialog.component';
import { MaterialModule } from './material/material.module'
import { CommonModule } from '@angular/common';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
@NgModule({
  declarations: [
    AppComponent,
    AdminComponent,
    ProjectRequestComponent,
    AddDialogComponent
  ],
  imports: [
    BrowserModule,
    MaterialModule,
    FormsModule,
    CommonModule,
    BrowserAnimationsModule,
    AppRoutingModule
  ],
  providers: [AddDialogComponent],
  bootstrap: [AppComponent, ProjectRequestComponent, AdminComponent]
})
export class AppModule { }
