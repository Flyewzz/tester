import React from "react";
import "./Rules.css";

function Rules(props) {
  const cases = [
    {
      code: "OK",
      transcript: "OK",
      description: "Тест пройден",
    },
    {
      code: "WA",
      transcript: "Wrong answer",
      description: "Программа выдает неправильный результат",
    },
    {
      code: "CE",
      transcript: "Compilation error",
      description: `Произошла ошибка компиляции (указывается сообщение об ошибке в поле
        сообщений)`,
    },
    {
      code: "TL",
      transcript: "Time limit",
      description: `Программа работает слишком долго или неоптимально`,
    },
    {
      code: "ML",
      transcript: "Memory limit",
      description: `Программа расходует слишком много оперативной памяти (например, при бесконечной рекурсии)`,
    },
  ];
  const getCases = (cases) => {
    return cases.map((state) => {
      return (
        <tr>
          <td>
            <i>
              <b
                className={
                  state.code === "OK" ? "positive-answer" : "negative-answer"
                }
              >
                {state.code}
              </b>
            </i>
          </td>
          <td>{state.transcript}</td>
          <td>{state.description}</td>
        </tr>
      );
    });
  };

  return (
    <table className="table-rules">
      <caption style={{ marginBottom: "10px" }}>Условные обозначения</caption>
      <thead>
        <th>Код возврата</th>
        <th>Расшифровка</th>
        <th>Значение</th>
      </thead>
      <tbody>{getCases(cases)}</tbody>
    </table>
  );
}

export default Rules;
