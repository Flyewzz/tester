import React from "react";
import ReactQuill from "react-quill";
import {
  Card,
  CardBody,
  Form,
  FormInput,
  CardColumns,
  CardTitle,
} from "shards-react";

import "react-quill/dist/quill.snow.css";
import "../../assets/quill.css";

const Test = (props) => (
  <Card small className="mb-3">
    <CardTitle style={{textAlign: "center", paddingTop: "10px"}}>
      <strong>Test #01</strong>
    </CardTitle>
    <CardBody>
      <Form className="add-new-post">
        <CardColumns>
          <textarea
            className="ql-container ql-snow"
            style={{ height: "100%", width: "100%" }}
          ></textarea>
          <textarea
            className="ql-container ql-snow"
            style={{ height: "100%" }}
          ></textarea>
        </CardColumns>
      </Form>
    </CardBody>
  </Card>
);

export default Test;
