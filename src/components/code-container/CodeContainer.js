import React from "react";
import "./CodeContainer.css";
import CodeMirror from "@uiw/react-codemirror";
import "codemirror/addon/display/autorefresh";
import "codemirror/addon/comment/comment";
import "codemirror/addon/edit/matchbrackets";
import "codemirror/keymap/sublime";
import "codemirror/theme/monokai.css";
import AceEditor from "react-ace";

import "ace-builds/src-noconflict/mode-c_cpp";
import "ace-builds/src-noconflict/theme-github";

function CodeContainer(props) {
  let modeNames = {
    "python": "text/x-python",
    "c++": "text/x-c++src",
  };
  return (
    <div id="container">
      <AceEditor
        mode="c_cpp"
        theme="github"
        onChange={props.handleChange}
        name="UNIQUE_ID_OF_DIV"
        setOptions={{
          enableBasicAutocompletion: true,
          enableLiveAutocompletion: true,
          enableSnippets: true,
        }}
        style={{ height: '400px' }}
        ref={instance => { this.ace = instance; }} // Let's put things into scope
      />
      <form
        encType="multipart/form-data"
      >
        <input
          className={"input"}
          type="submit"
          name="preview-form-submit"
          id="preview-form-submit"
          value="Submit"
          onClick={props.onClick}
        />
      </form>

      {/* var editor = CodeMirror.fromTextArea(document.getElementById("code"), {
          lineNumbers: true, // Нумеровать каждую строчку.
          matchBrackets: true,
          // mode: "text/x-c++src",
          mode: "text/x-python",
          indentUnit: 2, // Длина отступа в пробелах.
          indentWithTabs: true,
          enterMode: "keep",
          tabMode: "shift",
          theme: "monokai" */}
    </div>
  );
}

export default CodeContainer;
