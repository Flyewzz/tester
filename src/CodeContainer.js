import React from "react";
import "./CodeContainer.css";
import CodeMirror from "@uiw/react-codemirror";
import "codemirror/addon/display/autorefresh";
import "codemirror/addon/comment/comment";
import "codemirror/addon/edit/matchbrackets";
import "codemirror/keymap/sublime";
import "codemirror/theme/monokai.css";

function CodeContainer(props) {
  let modeNames = {
    'python': 'text/x-python',
    'c++': 'text/x-c++src',
  }
  return (
    <div id="container">
      <CodeMirror
        value={props.code}
        options={{
          theme: "monokai",
          tabSize: 2,
          keyMap: "sublime",
          lineNumbers: true, // Нумеровать каждую строчку.
          matchBrackets: true,
          mode: modeNames[props.lang],
          indentUnit: 2, // Длина отступа в пробелах.
          indentWithTabs: true,
          enterMode: "keep",
          tabMode: "shift",
        }}
        height="400px"
        className={"CodeMirror"}
        onChange={props.onChange}
      />
      {/* <textarea
        id="code"
        cols="60"
        rows="7"
        onChange={props.onChange}
      ></textarea> */}
      <form
        encType="multipart/form-data"
        // id="upload-container"
        // method="POST"
        // action="/test"
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
