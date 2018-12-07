import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { mergeMap, map } from 'rxjs/operators';
const BASE_URL = 'https://deployment-function-holder.appspot.com';

@Component({
  selector: 'app-view',
  templateUrl: './view.component.html',
  styleUrls: ['./view.component.css']
})
export class ViewComponent implements OnInit {

  project_request: any;
  /**
   * Component constructor
   */
  constructor() {

  }

  /**
   * OnInit implementation
   */
  ngOnInit() {
    
  }

  
}
