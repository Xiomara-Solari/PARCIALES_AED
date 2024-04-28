FRR|F|12345|NOMBREAPEL# + FRM|F|12345|NOMBREAPEL# 

EE|NN|----FDS

ACCION recu2023 ES
	AMBIENTE
		docentes,jubilar: secuencia de caracteres;
		d:caracter;
		edad: secuencia de enteros;
		e:entero;

		region_us: caracter;
		tot_doc, tot_nodoc: entero;

		PROCEDIMIENTO copiar_salida() ES
			MIENTRAS (d<>"#") HACER:
				GRABAR(jubilar,d);
				AVZ(docentes,d);
			Finmientras;
			tot_doc:=tot_doc +1;
		Finproced;

		PROCEDIMIENTO tratar_aporte() ES
			//consideramos que por ventana hay un digito
			num1:= e; //primer caracter de la edad
			AVZ(edad,e);
			num2:= e;
			eda:= num1*10 + num2;
			AVZ(edad,e);
			//primer car de el aporte
			a1:= e;
			AVZ(edad,e);
			a2:= e;
			aporte:= a1*10 + a2;
			AVZ(edad,e); //en la siguiente persona
		Finproced;

		PROCEDIMIENTO avanzar() ES	
			REPETIR
				AVZ(docentes,d);
			HASTA QUE (v="#")
		Finproced;

	PROCESO
		ARR(docentes);
		AVZ(docentes,d);
		ARR(edad);
		AVZ(edad,e);
		CREAR(jubilar);

		tot_doc:=0;
		tot_nodoc:=0;

		Esc("Ingrese una regional:");
		Leer(region_us); //suponemos que el ususario sabe cuáles son las regionales para ingresar

		MIENTRAS NFDS(docentes) HACER:
			MIENTRAS (d<>"+") HACER: //el + separa las regionales
				AVZ(docentes,d);
				AVZ(docentes,d);
				resg_reg:=d; //resguardo la region
				MIENTRAS (d<>"#") HACER:
					SI (resg_reg = region_us) ENTONCES: //region coincide con la del usuario
						AVZ(docentes,d); //avanza la regio
						res_sex:= d;
						AVZ(docentes,d);
						//aca estoy en el legajo en sec de docentes

						tratar_aporte()
						SI (aporte >= 30) ENTONCES: //verifico que pueda jubilarse y copio
							SI (res_sex = "M") ENTONCES:
								SI (eda >=65) ENTONCES:
									copiar_salida();
								fins;
							SINO
								SI (eda >=60) ENTONCES:
									copiar_salida();	
								fins;
							fins;
						SINO
							tot_nodoc:=tot_nodoc +1 //no se pueden jubilar
						finsi

					SINO //trata cuando la region es distinta a la des usuario
						tratar_aporte();
						SI (aporte >= 30) ENTONCES: 
							SI (res_sex = "M") ENTONCES:
								SI (eda >=65) ENTONCES:
									tot_doc:=tot_doc +1	
								fins;
							SINO
								SI (eda >=60) ENTONCES:
									tot_doc:=tot_doc +1	
								fins;
							fins;
						SINO
							tot_nodoc:=tot_nodoc +1
						Fins;
						avanzar(); //avanzo hasta el "#"
					Fins;	
				Finmientras;
				AVZ(docentes,d);
			Finmientras;
			AVZ(docentes,d);

			Esc("El total de docentes que pueden jubilarse de la regionl",resg_reg,"es:",tot_doc);
			Esc("El total de docentes que no pueden jubilarse de la regionl",resg_reg,"es:",tot_nodoc);
		Finmientras;
		CERRAR(docentes);
		CERRAR(edad);
		CERRAR(jubilar);
Finaccion


ACCION recu2023_2 ES
	AMBIENTE
		JUBILAR = Registro
			regional: AN(30);
			carrera: AN(30);
			legajo: N(5);
			sexo: {"m","f"};
			nom: AN(30);
		Finregistro;

		Arch: archivo de JUBILAR ordenado por regional, carrera y legajo;
		reg: JUBILAR;

		JUBILAR_SAL = Registro
			regional: AN(30);
			cant_jub: N(10);
		Finregistro;

		Arch_sal: archivo de JUBILAR_SAL;
		reg_sal: JUBILAR_SAL;

		resg_reg: AN(30);
		resg_car: AN(30);
		resg_regmay: AN(30);

		tot_reg, tot_carf,tot_carm,tot_gral,tot_mayor: entero;

		PROCEDIMIENTO corte_regional() ES	
			corte_carrera();
			Esc("En la regional",resg_reg,", la cantidad de docentes es:",tot_reg);
			tot_gral:=tot_gral+tot_reg;

			reg_sal.regional:= resg_reg;
			reg_sal.cant_jub:= tot_reg;
			GRABAR(Arch_sal,reg_sal);
			
			SI (tot_reg < tot_mayor) ENTONCES	
				tot_mayor:= tot_reg;
				resg_regmay:= reg.regional;
			Finsi;

			tot_reg:=0;
			resg_reg:=reg.regional;
		Finproced;

		PROCEDIMIENTO corte_carrera() ES
			Esc("El total de docentes a jubilarse de sexo femenino es:",tot_carf);
			Esc("El total de docentes a jubilarse de sexo masculino es:",tot_carm);
			tot_reg:= tot_reg + tot_carf + tot_carm;
			tot_carf:=0;
			tot_carm:=0;
			resg_car:= reg.carrera;
		Finproced;

	PROCESO	
		ABRIR E/(Arch);
		LEER(Arch,reg);
		ABRIR S/(Arch_sal);

		tot_reg:= 0;
		tot_carf:=0;
		tot_carm:=0;
		tot_gral:=0;
		tot_mayor:=0;

		resg_reg:= reg.regional;
		resg_car:= reg.carrera;
		resg_regmay:= reg.regional;

		MIENTRAS NFDA(Arch) HACER:
			SI (resg_reg <> reg.regional) ENTONCES
				corte_regional();
			SINO
				SI (resg_car <> reg.carrera) ENTONCES:
					corte_carrera();
				Finsi;
			Finsi;

			SI (reg.sexo ="m") ENTONCES
				tot_carm:= tot_carm + 1;
			SINO
				tot_carf:= tot_carf + 1;
			Finsi;

			LEER(Arch,reg);
		Finmientras;
		corte_regional;
		CERRAR(Arch);
		CERRAR(Arch_sal);
		Esc("El total general de docentes a jubilar es:",tot_gral);
		Esc("La regional",resg_regmay,"presenta la mayor cantidad de docentes a jubilar y es de:"tot_mayor);
Finaccion.
