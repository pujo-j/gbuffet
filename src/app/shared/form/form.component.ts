import { Component, OnInit, Output, Input, EventEmitter, OnChanges } from '@angular/core';

@Component({
  selector: 'app-form',
  templateUrl: './form.component.html',
  styleUrls: ['./form.component.css']
})
export class FormComponent implements OnInit , OnChanges {
  @Input() model: any;
  // Admin update the form to response to the project request user
  isUpdateMode: boolean = false;
  isAdminMode: boolean = false;

  @Output('cancel') cancel$: EventEmitter<any>;
  @Output('submit') submit$: EventEmitter<any>;

  constructor() {
    this.submit$ = new EventEmitter();
    this.cancel$ = new EventEmitter();
    this.model = { allocation:{type: "compute"} };
  }

  /**
   * OnInit implementation
   */
  ngOnInit() {
    console.log('update: ', this.isUpdateMode);
  }

  /**
   * Function to handle component update
   *
   * @param record
   */
  ngOnChanges(record) {
    if (record.model && record.model.currentValue) {
      this.model = record.model.currentValue;
      this.isUpdateMode = !!this.model;
    }
  }
  cancel() {
    this.cancel$.emit();
  }

  submit(project: any) {
    console.log(this.model)
    this.submit$.emit(this.model);
  }
  
}
