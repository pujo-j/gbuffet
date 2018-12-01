import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';


import {
  MatChipsModule,
  MatToolbarModule,
  MatCardModule,
  MatTabsModule,
  MatButtonModule,
  MatInputModule,
  MatCheckboxModule,
  MatRadioModule,
  MatIconModule,
  MatListModule,
  MatDialogModule
} from "@angular/material";

let materialModules = [
  MatChipsModule,
  MatToolbarModule,
  MatCardModule,
  MatTabsModule,
  MatButtonModule,
  MatInputModule,
  MatCheckboxModule,
  MatRadioModule,
  MatIconModule,
  MatListModule,
  MatDialogModule
];
@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    materialModules
  ],
  exports:[
    materialModules
  ]
})


export class MaterialModule { }
