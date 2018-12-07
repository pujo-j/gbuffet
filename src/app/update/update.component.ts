import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { mergeMap, map } from 'rxjs/operators';
const BASE_URL = 'https://deployment-function-holder.appspot.com';

@Component({
  selector: 'app-update',
  templateUrl: './update.component.html',
  styleUrls: ['./update.component.css']
})
export class UpdateComponent implements OnInit {
  project_request: any;
  /**
   * Component constructor
   */
  constructor(private _route: ActivatedRoute, private _router: Router, private _http: HttpClient) {
    this.project_request = {
    };
  }

  /**
   * OnInit implementation
   */
  ngOnInit() {
    this._route.params
      .pipe(
        map((params: any) => params.id),
        mergeMap((id: string) => this._http.get(`${BASE_URL}/project_requests/${id}`))
      )
      .subscribe((project_request: any) => (this.project_request = project_request));
  }

}
