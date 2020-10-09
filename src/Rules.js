import React from "react";
import './Rules.css';


function Rules(props) {
  return (
    <table className='table-rules'>
      <caption style={{ marginBottom: "10px" }}>Условные обозначения</caption>
      <thead>
        <th>Код возврата</th>
        <th>Расшифровка</th>
        <th>Значение</th>
      </thead>
      <tbody>
        <tr>
          <td>
            <i>
              <b>OK</b>
            </i>
          </td>
          <td>OK</td>
          <td>Тест пройден</td>
        </tr>
        <tr>
          <td>
            <i>
              <b>WA</b>
            </i>
          </td>
          <td>Wrong Answer</td>
          <td>Программа выдает неправильный результат</td>
        </tr>
        <tr>
          <td>
            <i>
              <b>CE</b>
            </i>
          </td>
          <td>Compilation Error</td>
          <td>
            Произошла ошибка компиляции (указывается сообщение об ошибке в поле
            сообщений)
          </td>
        </tr>
        <tr>
          <td>
            <i>
              <b>TL</b>
            </i>
          </td>
          <td>Time Limit</td>
          <td>Программа работает слишком долго или неоптимально</td>
        </tr>
        <tr>
          <td>
            <i>
              <b>ML</b>
            </i>
          </td>
          <td>Memory Limit</td>
          <td>Программа расходует слишком много оперативной памяти (например, при бесконечной рекурсии)</td>
        </tr>
      </tbody>
    </table>
  );
}

export default Rules;
