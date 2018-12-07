import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { AdminComponent } from './admin';
import { ProjectRequestComponent } from './project-request';
import { UpdateComponent } from './update';
import { ViewComponent } from './view';

const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'home', component: ProjectRequestComponent },
  { path: 'admin', component: AdminComponent },
  { path: 'edit/:id', component: UpdateComponent },
  { path: 'view/:id', component: ViewComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
