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
import { HttpClientModule, HttpClient } from '@angular/common/http';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import { FormComponent } from './shared/form/form.component';
import { CardComponent } from './shared/card/card.component';
@NgModule({
  declarations: [
    AppComponent,
    AdminComponent,
    ProjectRequestComponent,
    AddDialogComponent,
    FormComponent,
    CardComponent
  ],
  imports: [
    BrowserModule,
    MaterialModule,
    FormsModule,
    CommonModule,
    BrowserAnimationsModule,
    AppRoutingModule
  ],
  entryComponents: [AddDialogComponent],
  providers: [HttpClient],
  bootstrap: [AppComponent, ProjectRequestComponent, AdminComponent]
})
export class AppModule { }
