import React from "react";

function Rules(props) {
  return (
    <table>
      <caption style={{ marginBottom: "10px" }}>Условные обозначения</caption>
      <thead>
        <th>Код возврата</th>
        <th>Расшифровка</th>
        <th>Значение</th>
      </thead>
      <tbody>
        <tr>
          <td style={{ color: "green", textAlign: "center" }}>
            <i>
              <b>OK</b>
            </i>
          </td>
          <td>OK</td>
          <td>Тест пройден</td>
        </tr>
        <tr>
          <td style={{ color: "red", textAlign: "center" }}>
            <i>
              <b>WA</b>
            </i>
          </td>
          <td>Wrong Answer</td>
          <td>Программа выдает неправильный результат</td>
        </tr>
        <tr>
          <td style={{ color: "red", textAlign: "center" }}>
            <i>
              <b>CE</b>
            </i>
          </td>
          <td>Compilation Error</td>
          <td>
            {" "}
            Произошла ошибка компиляции (указывается сообщение об ошибке в поле
            сообщений)
          </td>
        </tr>
        <tr>
          <td style={{ color: "red", textAlign: "center" }}>
            <i>
              <b>TL</b>
            </i>
          </td>
          <td>Time Limit</td>
          <td>Программа выдает неправильный результат</td>
        </tr>
      </tbody>
    </table>
  );
}

export default Rules;
