import React from "react";
import PageTitle from "../components/common/PageTitle";
import Editor from "../components/add-new-post/Editor";
import SidebarActions from "../components/add-new-post/SidebarActions";
import SidebarCategories from "../components/add-new-post/SidebarCategories";
import { Container, Row, Col, ListGroup, ListGroupItem } from "shards-react";
import CustomFileUpload from "../components/components-overview/CustomFileUpload";
import { useState } from "react";
import Test from "../components/add-new-post/Test";

const AddNewPost = () => {
  const [state, setState] = useState({
    taskTitle: "",
    taskText: "",
  });

  const handleChange = (e) => {
    const value = e.target ? e.target.value : e;
    setState({
      ...state,
      taskText: value,
    });
    // switch (e.target.name) {
    // }
  };

  return (
    <Container fluid className="main-content-container px-4 pb-4">
      {/* Page Header */}
      <Row noGutters className="page-header py-4">
        <PageTitle
          sm="4"
          title="Add New Post"
          subtitle="Blog Posts"
          className="text-sm-left"
        />
      </Row>

      <Row>
        {/* Editor */}
        <Col lg="9" md="12">
          <Editor onChange={handleChange} text={state.taskText} />
        </Col>

        {/* Sidebar Widgets */}
        <Col lg="3" md="12">
          <SidebarActions />
          <ListGroupItem className="px-3">
            <strong className="text-muted d-block mb-2">
              Custom File Upload
            </strong>
            <CustomFileUpload />
          </ListGroupItem>
          {/* <SidebarCategories /> */}
        </Col>
      </Row>
      <Col lg="12" lg="4">
        <Test />
      </Col>
      <Row></Row>
    </Container>
  );
};

export default AddNewPost;
