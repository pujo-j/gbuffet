import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AdminComponent } from './admin/admin.component';
import { ProjectRequestComponent } from './project-request/project-request.component';
import { AddDialogComponent } from './project-request/add-dialog/add-dialog.component';
import { MaterialModule } from './material/material.module'
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
    AppRoutingModule
  ],
  providers: [AddDialogComponent],
  bootstrap: [AppComponent, ProjectRequestComponent, AdminComponent]
})
export class AppModule { }
